//
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
//

package object

import (
	"github.com/insolar/insolar/instrumentation/insmetrics"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	inmemoryStorage = insmetrics.MustTagKey("inmemorystorage")
	dbStorage       = insmetrics.MustTagKey("dbstorage")
)

var (
	statLifelineInMemoryAddedCount = stats.Int64(
		"object/lifeline/added/count",
		"How many index-records have been saved in in-indexStorage index storage",
		stats.UnitDimensionless,
	)
	statLifelineInMemoryRemovedCount = stats.Int64(
		"object/lifeline/removed/count",
		"How many index-records have been removed from an index storage",
		stats.UnitDimensionless,
	)
	statObjectPendingRequestsInMemoryAddedCount = stats.Int64(
		"оbject/pendings/added/count",
		"How many index-records have been saved in in-indexStorage index storage",
		stats.UnitDimensionless,
	)
	statObjectPendingRequestsInMemoryRemovedCount = stats.Int64(
		"оbject/pendings/removed/count",
		"How many index-records have been saved in in-indexStorage index storage",
		stats.UnitDimensionless,
	)
	statIndexDBAddedCount = stats.Int64(
		"object/lifeline/added/count",
		"How many index-records have been saved in in-indexStorage index storage",
		stats.UnitDimensionless,
	)
	statRecordInMemoryAddedCount = stats.Int64(
		"recordstorage/added/count",
		"How many records have been saved to a in-memory storage",
		stats.UnitDimensionless,
	)
	statRecordInMemoryRemovedCount = stats.Int64(
		"recordstorage/removed/count",
		"How many records have been removed from a in-memory storage",
		stats.UnitDimensionless,
	)
)

func init() {
	err := view.Register(
		&view.View{
			Name:        statLifelineInMemoryAddedCount.Name(),
			Description: statLifelineInMemoryAddedCount.Description(),
			Measure:     statLifelineInMemoryAddedCount,
			Aggregation: view.Count(),
			TagKeys:     []tag.Key{inmemoryStorage},
		},
		&view.View{
			Name:        statLifelineInMemoryRemovedCount.Name(),
			Description: statLifelineInMemoryRemovedCount.Description(),
			Measure:     statLifelineInMemoryRemovedCount,
			Aggregation: view.Count(),
			TagKeys:     []tag.Key{inmemoryStorage},
		},
		&view.View{
			Name:        statIndexDBAddedCount.Name(),
			Description: statIndexDBAddedCount.Description(),
			Measure:     statIndexDBAddedCount,
			Aggregation: view.Count(),
			TagKeys:     []tag.Key{dbStorage},
		},
		&view.View{
			Name:        statRecordInMemoryAddedCount.Name(),
			Description: statRecordInMemoryAddedCount.Description(),
			Measure:     statRecordInMemoryAddedCount,
			Aggregation: view.Count(),
			TagKeys:     []tag.Key{inmemoryStorage},
		},
		&view.View{
			Name:        statRecordInMemoryRemovedCount.Name(),
			Description: statRecordInMemoryRemovedCount.Description(),
			Measure:     statRecordInMemoryRemovedCount,
			Aggregation: view.Count(),
			TagKeys:     []tag.Key{inmemoryStorage},
		},
	)
	if err != nil {
		panic(err)
	}
}
