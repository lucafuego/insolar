// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/insolar/blob/master/LICENSE.md.

package main

import (
	"github.com/insolar/insolar/application/genesisrefs"
	"github.com/pkg/errors"
)

func getAPIInfoResponse() (map[string]interface{}, error) {
	rootDomain := genesisrefs.ContractRootDomain
	if rootDomain.IsEmpty() {
		return nil, errors.New("rootDomain ref is nil")
	}

	rootMember := genesisrefs.ContractRootMember
	if rootMember.IsEmpty() {
		return nil, errors.New("rootMember ref is nil")
	}

	migrationDaemonMembers := genesisrefs.ContractMigrationDaemonMembers
	migrationDaemonMembersStrs := make([]string, 0)
	for _, r := range migrationDaemonMembers {
		if r.IsEmpty() {
			return nil, errors.New("migration daemon members refs are nil")
		}
		migrationDaemonMembersStrs = append(migrationDaemonMembersStrs, r.String())
	}

	migrationAdminMember := genesisrefs.ContractMigrationAdminMember
	if migrationAdminMember.IsEmpty() {
		return nil, errors.New("migration admin member ref is nil")
	}
	feeMember := genesisrefs.ContractFeeMember
	if feeMember.IsEmpty() {
		return nil, errors.New("feeMember ref is nil")
	}
	return map[string]interface{}{
		"rootDomain":             rootDomain.String(),
		"rootMember":             rootMember.String(),
		"migrationAdminMember":   migrationAdminMember.String(),
		"feeMember":              feeMember.String(),
		"migrationDaemonMembers": migrationDaemonMembersStrs,
	}, nil
}
