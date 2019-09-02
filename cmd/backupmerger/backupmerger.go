///
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
///

package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/jet"
	"github.com/insolar/insolar/insolar/pulse"
	"github.com/insolar/insolar/insolar/store"
	"github.com/insolar/insolar/ledger/heavy/executor"
	"github.com/insolar/insolar/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func usage() {
	pflag.Usage()
	os.Exit(0)
}

func closeRawDB(bdb *badger.DB, err error) {
	closeError := bdb.Close()
	if closeError != nil || err != nil {
		printError("failed to close db", closeError)
		printError("", err)
		os.Exit(1)
	}
}

func merge(targetDBPath string, backupFileName string, numberOfWorkers int) {
	ops := badger.DefaultOptions(targetDBPath)
	bdb, err := badger.Open(ops)
	if err != nil {
		printError("failed to open badger", err)
		os.Exit(1)
	}

	if err := isDBEmpty(bdb); err == nil {
		closeRawDB(bdb, errors.New("db must not be empty"))
		return
	}
	log.Info("DB is not empty")

	bkpFile, err := os.Open(backupFileName)
	if err != nil {
		closeRawDB(bdb, err)
		return
	}

	err = bdb.Load(bkpFile, numberOfWorkers)
	if err != nil {
		closeRawDB(bdb, err)
		return
	}
	log.Info("Successfully merged")
	closeRawDB(bdb, nil)
}

func parseMergeParams() *cobra.Command {
	var (
		targetDBPath    string
		backupFileName  string
		numberOfWorkers int
	)

	var mergeCmd = &cobra.Command{
		Use:   "merge",
		Short: "merge incremental backup to existing db",
		Run: func(cmd *cobra.Command, args []string) {
			merge(targetDBPath, backupFileName, numberOfWorkers)
		},
	}
	mergeFlags := mergeCmd.Flags()
	targetDBFlagName := "target-db"
	bkpFileName := "bkp-name"
	mergeFlags.StringVarP(
		&targetDBPath, targetDBFlagName, "t", "", "directory where backup will be roll to (required)")
	mergeFlags.StringVarP(
		&backupFileName, bkpFileName, "n", "", "file name if incremental backup (required)")
	mergeFlags.IntVarP(
		&numberOfWorkers, "workers-num", "w", 1, "number of workers to read backup file")

	cobra.MarkFlagRequired(mergeFlags, targetDBFlagName)
	cobra.MarkFlagRequired(mergeFlags, bkpFileName)

	return mergeCmd
}

type dbInitializedKey insolar.PulseNumber

func (k dbInitializedKey) Scope() store.Scope {
	return store.ScopeDBInit
}

func (k dbInitializedKey) ID() []byte {
	bytes, err := time.Now().MarshalBinary()
	if err != nil {
		panic("failed to marshal time: " + err.Error())
	}
	return bytes
}

func isDBEmpty(bdb *badger.DB) error {
	tableInfo := bdb.Tables(true)
	if len(tableInfo) != 0 {
		return errors.New("tableInfo is not empty")
	}

	lsm, vlog := bdb.Size()
	if lsm != 0 || vlog != 0 {
		println("lsm: ", lsm, ", vlog: ", vlog)
		return errors.New("lsm ot vlog are not empty")
	}

	return nil
}

func createEmptyBadger(dbDir string) {
	ops := badger.DefaultOptions(dbDir)
	var err error
	bdb, err := badger.Open(ops)
	if err != nil {
		printError("failed to open badger", err)
		os.Exit(1)
	}

	err = isDBEmpty(bdb)
	if err != nil {
		closeRawDB(bdb, err)
		return
	}
	log.Info("DB is empty")

	var key dbInitializedKey
	fullKey := append(key.Scope().Bytes(), key.ID()...)

	err = bdb.Update(func(txn *badger.Txn) error {
		return txn.Set(fullKey, []byte{})
	})

	t := time.Time{}
	t.UnmarshalBinary(key.ID())
	log.Info("Set db initialized key: ", t.String())

	if err != nil {
		closeRawDB(bdb, err)
		return
	}

	closeRawDB(bdb, nil)
}

func parseCreateParams() *cobra.Command {
	var dbDir string
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "create new empty badger",
		Run: func(cmd *cobra.Command, args []string) {
			createEmptyBadger(dbDir)
		},
	}

	dbDirFlagName := "db-dir"
	createCmd.Flags().StringVarP(
		&dbDir, dbDirFlagName, "d", "", "directory where new badger will be created (required)")

	cobra.MarkFlagRequired(createCmd.Flags(), dbDirFlagName)

	return createCmd
}

func prepareBackup(dbDir string) {
	// finalizing of data
	ops := badger.DefaultOptions(dbDir)
	bdb, err := store.NewBadgerDB(ops)
	if err != nil {
		printError("failed to open badger", err)
		os.Exit(1)
	}
	ctx := context.Background()
	closeDB := func(err error) {
		errStop := bdb.Stop(ctx)
		if err != nil || errStop != nil {
			printError("", err)
			printError("failed to close db", errStop)
			os.Exit(1)
		}
	}

	pulsesDB := pulse.NewDB(bdb)

	jetKeeper := executor.NewJetKeeper(jet.NewDBStore(bdb), bdb, pulsesDB)
	log.Info("Current top sync pulse: ", jetKeeper.TopSyncPulse().String())

	it := bdb.NewIterator(executor.BackupStartKey(math.MaxUint32), true)
	if !it.Next() {
		closeDB(errors.New("no backup start keys"))
		return
	}

	pulseNumber := insolar.NewPulseNumber(it.Key())
	log.Info("Found backup start key: ", pulseNumber.String())

	if !jetKeeper.HasAllJetConfirms(ctx, pulseNumber) {
		closeDB(errors.New("data is inconsistent. pulse " + pulseNumber.String() + " must has all confirms"))
		return
	}

	log.Info("All jet confirmed for pulse: ", pulseNumber.String())
	err = jetKeeper.AddBackupConfirmation(ctx, pulseNumber)
	if err != nil {
		closeDB(errors.New("failed to add backup confirmation for pulse" + pulseNumber.String()))
		return
	}

	if jetKeeper.TopSyncPulse() != pulseNumber {
		closeDB(errors.New("new top sync pulse must be equal to last backuped"))
		return
	}

	log.Info("New top sync pulse: ", jetKeeper.TopSyncPulse().String())
	closeDB(nil)
}

func parsePrepareBackupParams() *cobra.Command {
	var dbDir string
	var prepareBackupCmd = &cobra.Command{
		Use:   "prepare_backup",
		Short: "prepare backup for usage",
		Run: func(cmd *cobra.Command, args []string) {
			prepareBackup(dbDir)
		},
	}

	dbDirFlagName := "db-dir"
	prepareBackupCmd.Flags().StringVarP(
		&dbDir, dbDirFlagName, "d", "", "directory where new badger will be created (required)")

	cobra.MarkFlagRequired(prepareBackupCmd.Flags(), dbDirFlagName)

	return prepareBackupCmd
}

func parseInputParams() {

	var rootCmd = &cobra.Command{
		Use:   "backupmanager",
		Short: "backupmanager is the command line client for managing backups",
	}

	rootCmd.AddCommand(parseMergeParams())
	rootCmd.AddCommand(parseCreateParams())
	rootCmd.AddCommand(parsePrepareBackupParams())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printError(message string, err error) {
	if err == nil {
		return
	}
	println(errors.Wrap(err, "ERROR "+message).Error())
}

func main() {
	log.SetLevel("Debug")

	cfg := configuration.NewLog()
	cfg.Level = "Debug"
	cfg.Formatter = "text"
	l, _ := log.NewLog(cfg)

	log.SetGlobalLogger(l)
	parseInputParams()
}
