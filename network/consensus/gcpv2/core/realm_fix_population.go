//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package core

import (
	"context"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network/consensus/common/consensuskit"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/census"
)

func NewFixedRealmPopulation(strategy RoundStrategy, population census.OnlinePopulation, phase2ExtLimit uint8,
	fn NodeInitFunc) *FixedRealmPopulation {

	nodeCount := population.GetCount()
	otherCount := nodeCount
	if otherCount > 0 && !population.GetLocalProfile().IsJoiner() {
		otherCount-- //remove self when it is not a joiner
	}

	r := &FixedRealmPopulation{
		population: population,
		dynPop: dynPop{DynamicRealmPopulation{
			nodeInit:       fn,
			shuffleFunc:    strategy.ShuffleNodeSequence,
			baselineWeight: strategy.GetBaselineWeightForNeighbours(),
			phase2ExtLimit: phase2ExtLimit,
			indexedCount:   nodeCount,
			nodeIndex:      make([]*NodeAppearance, nodeCount),
			nodeShuffle:    make([]*NodeAppearance, otherCount),
			shuffledCount:  otherCount,
			dynamicNodes:   make(map[insolar.ShortNodeID]*NodeAppearance),
			indexedLenSet:  true, //locks down SealIndex
		}},
		bftMajorityCount: consensuskit.BftMajority(nodeCount),
	}
	r.initPopulation()
	ShuffleNodeAppearances(strategy.ShuffleNodeSequence, r.nodeShuffle)

	return r
}

var _ RealmPopulation = &FixedRealmPopulation{}

type dynPop struct{ DynamicRealmPopulation }

type FixedRealmPopulation struct {
	dynPop
	population census.OnlinePopulation

	bftMajorityCount int
}

func (r *FixedRealmPopulation) GetSealedLimit() (int, bool) {
	return len(r.nodeIndex), true
}

func (r *FixedRealmPopulation) initPopulation() {
	activeProfiles := r.population.GetProfiles()
	thisNodeID := r.population.GetLocalProfile().GetNodeID()

	nodes := make([]NodeAppearance, r.indexedCount)

	var j = 0
	for i, p := range activeProfiles {
		n := &nodes[i]
		r.nodeIndex[i] = n

		if p.GetOpMode().IsEvicted() {
			panic("illegal state")
		}

		n.init(p, nil, r.baselineWeight, r.CreatePacketLimiter())
		r.nodeInit(context.Background(), n)

		if p.GetNodeID() == thisNodeID {
			if r.self != nil {
				panic("schizophrenia")
			}
			r.self = n
		} else {
			r.nodeShuffle[j] = n
			j++
		}
	}
	if r.self == nil {
		panic("illegal state")
	}
}

func (r *FixedRealmPopulation) GetIndexedCount() int {
	return r.indexedCount
}

func (r *FixedRealmPopulation) GetBftMajorityCount() int {
	return r.bftMajorityCount
}

func (r *FixedRealmPopulation) GetActiveNodeAppearance(id insolar.ShortNodeID) *NodeAppearance {
	np := r.population.FindProfile(id)
	if np != nil && !np.IsJoiner() {
		return r.GetNodeAppearanceByIndex(np.GetIndex().AsInt())
	}
	return nil
}

func (r *FixedRealmPopulation) GetNodeAppearance(id insolar.ShortNodeID) *NodeAppearance {
	na := r.GetActiveNodeAppearance(id)
	if na != nil {
		return na
	}
	return r.GetJoinerNodeAppearance(id)
}

func (r *FixedRealmPopulation) GetNodeAppearanceByIndex(idx int) *NodeAppearance {
	return r.nodeIndex[idx]
}

func (r *FixedRealmPopulation) GetShuffledOtherNodes() []*NodeAppearance {
	return r.nodeShuffle
}

func (r *FixedRealmPopulation) GetIndexedNodes() []*NodeAppearance {
	return r.nodeIndex
}

func (r *FixedRealmPopulation) GetIndexedNodesAndHasNil() ([]*NodeAppearance, bool) {
	return r.nodeIndex, true
}

func (r *FixedRealmPopulation) SealIndex(indexedCountLimit int) bool {
	return r.indexedCount == indexedCountLimit
}

func (r *FixedRealmPopulation) AddToDynamics(n *NodeAppearance) (*NodeAppearance, error) {
	if !n.profile.IsJoiner() {
		panic("illegal value")
	}
	return r.dynPop.AddToDynamics(n)
}

func (r *FixedRealmPopulation) CreateVectorHelper() *RealmVectorHelper {

	v := r.DynamicRealmPopulation.CreateVectorHelper()
	v.realmPopulation = r
	return v
}

func (r *FixedRealmPopulation) appendDynamicNodes(nodes []*NodeAppearance) []*NodeAppearance {

	r.rw.RLock()
	defer r.rw.RUnlock()

	index := len(nodes)
	nodes = append(nodes, make([]*NodeAppearance, len(r.dynamicNodes))...)
	for _, v := range r.dynamicNodes {
		nodes[index] = v
		index++
	}
	return nodes
}

func (r *FixedRealmPopulation) GetAnyNodes(includeIndexed bool, shuffle bool) []*NodeAppearance {

	var nodes []*NodeAppearance
	joinerCount := r.GetJoinersCount()

	if !shuffle {
		if includeIndexed {
			nodes = append(make([]*NodeAppearance, 0, r.indexedCount+joinerCount), r.nodeIndex...)
		}
		nodes = r.appendDynamicNodes(nodes)
		return nodes
	}

	if includeIndexed {
		nodes = append(make([]*NodeAppearance, 0, r.indexedCount+joinerCount), r.nodeShuffle...)
		before := len(nodes)

		if !r.self.IsJoiner() {
			nodes = append(nodes, r.self)
		}
		nodes = r.appendDynamicNodes(nodes)
		if len(nodes) > before+reshuffleTolerance {
			ShuffleNodeAppearances(r.shuffleFunc, nodes)
		}
	}
	return nodes
}
