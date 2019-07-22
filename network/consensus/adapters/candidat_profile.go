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

package adapters

import (
	"fmt"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network/consensus/common/cryptkit"
	"github.com/insolar/insolar/network/consensus/common/endpoints"
	"github.com/insolar/insolar/network/consensus/common/longbits"
	"github.com/insolar/insolar/network/consensus/common/pulse"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/member"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/power"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/profiles"
	"time"
)

func NewCandidateProfile(
	address string,
	ref insolar.Reference,
	id insolar.ShortNodeID,
	primaryRole member.PrimaryRole,
	specialRole member.SpecialRole,
) *CandidateProfile {
	return &CandidateProfile{
		n:           NewOutbound(address),
		id:          id,
		primaryRole: primaryRole,
		specialRole: specialRole,
		ref:         ref,
	}
}

//func NewCandidateProfileFromJoinClaim(joinClaim *packets.NodeJoinClaim, isDiscovery bool) *CandidateProfile {
//	specialRole := member.SpecialRoleNone
//	if isDiscovery {
//		specialRole = member.SpecialRoleDiscovery
//	}
//
//	primaryRole := member.PrimaryRoleNeutral
//	switch joinClaim.NodeRoleRecID {
//	case insolar.StaticRoleVirtual:
//		primaryRole = member.PrimaryRoleVirtual
//	case insolar.StaticRoleHeavyMaterial:
//		primaryRole = member.PrimaryRoleHeavyMaterial
//	case insolar.StaticRoleLightMaterial:
//		primaryRole = member.PrimaryRoleLightMaterial
//	default:
//	}
//
//	return &CandidateProfile{
//		n:           NewOutbound(joinClaim.NodeAddress.String()),
//		id:          joinClaim.ShortNodeID,
//		primaryRole: primaryRole,
//		specialRole: specialRole,
//		ref:         joinClaim.NodeRef,
//	}
//}

type CandidateProfile struct {
	n           endpoints.Outbound
	id          insolar.ShortNodeID
	primaryRole member.PrimaryRole
	specialRole member.SpecialRole
	ref         insolar.Reference
}

func (c *CandidateProfile) GetJoinerSignature() cryptkit.SignatureHolder {
	return nil
}

func (c *CandidateProfile) GetIssuedAtPulse() pulse.Number {
	return 0
}

func (c *CandidateProfile) GetIssuedAtTime() time.Time {
	return time.Now()
}

func (c *CandidateProfile) GetPowerLevels() member.PowerSet {
	return member.PowerSet{0, 0, 0, 0xFF}
}

func (c *CandidateProfile) GetExtraEndpoints() []endpoints.Outbound {
	return nil
}

func (c *CandidateProfile) GetIssuerID() insolar.ShortNodeID {
	return 0
}

func (c *CandidateProfile) GetIssuerSignature() cryptkit.SignatureHolder {
	return nil
}

func (c *CandidateProfile) GetNodePublicKey() cryptkit.SignatureKeyHolder {
	v := &longbits.Bits512{}
	longbits.FillBitsWithStaticNoise(uint32(c.id), v[:])
	k := cryptkit.NewSignatureKey(v, "stub/stub", cryptkit.PublicAsymmetricKey)
	return &k
}

func (c *CandidateProfile) GetStartPower() member.Power {
	return 10
}

func (c *CandidateProfile) GetReference() insolar.Reference {
	return c.ref
}

func (c *CandidateProfile) ConvertPowerRequest(request power.Request) member.Power {
	if ok, cl := request.AsCapacityLevel(); ok {
		return member.PowerOf(uint16(cl.DefaultPercent()))
	}
	_, pw := request.AsMemberPower()
	return pw
}

func (c *CandidateProfile) GetPrimaryRole() member.PrimaryRole {
	return c.primaryRole
}

func (c *CandidateProfile) GetSpecialRoles() member.SpecialRole {
	return c.specialRole
}

func (*CandidateProfile) IsAllowedPower(p member.Power) bool {
	return true
}

func (c *CandidateProfile) GetDefaultEndpoint() endpoints.Outbound {
	return c.n
}

func (*CandidateProfile) GetPublicKeyStore() cryptkit.PublicKeyStore {
	return nil
}

func (c *CandidateProfile) IsAcceptableHost(from endpoints.Inbound) bool {
	addr := c.n.GetNameAddress()
	return addr.Equals(from.GetNameAddress())
}

func (c *CandidateProfile) GetStaticNodeID() insolar.ShortNodeID {
	return c.id
}

func (c *CandidateProfile) GetIntroducedNodeID() insolar.ShortNodeID {
	return c.id
}

func (c *CandidateProfile) GetExtension() profiles.StaticProfileExtension {
	return c
}

func (c *CandidateProfile) String() string {
	return fmt.Sprintf("{sid:%v, n:%v}", c.id, c.n)
}
