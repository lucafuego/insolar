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

package gateway

import (
	"context"
	"github.com/insolar/insolar/log"
	"time"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/certificate"
	"github.com/insolar/insolar/component"
	"github.com/insolar/insolar/insolar/pulse"
	"github.com/insolar/insolar/instrumentation/inslogger" // TODO remove before merge
	"github.com/insolar/insolar/network/gateway/bootstrap"
	"github.com/insolar/insolar/network/hostnetwork/host"
	"github.com/insolar/insolar/network/hostnetwork/packet"
	"github.com/insolar/insolar/network/hostnetwork/packet/types"
	"github.com/insolar/insolar/network/storage"
	"github.com/insolar/insolar/platformpolicy"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/network"
)

// Base is abstract class for gateways

type Base struct {
	component.Initer

	Self                network.Gateway
	Gatewayer           network.Gatewayer           `inject:""`
	NodeKeeper          network.NodeKeeper          `inject:""`
	ContractRequester   insolar.ContractRequester   `inject:""`
	CryptographyService insolar.CryptographyService `inject:""`
	CertificateManager  insolar.CertificateManager  `inject:""`
	HostNetwork         network.HostNetwork         `inject:""`
	PulseAccessor       pulse.Accessor              `inject:""`
	PulseAppender       pulse.Appender              `inject:""`
	PulseManager        insolar.PulseManager        `inject:""`
	BootstrapRequester  bootstrap.Requester         `inject:""`
	Rules               network.Rules               `inject:""`
	KeyProcessor        insolar.KeyProcessor        `inject:""`

	// DiscoveryBootstrapper bootstrap.DiscoveryBootstrapper `inject:""`
	bootstrapETA insolar.PulseNumber

	candidateProfile packet.CandidateProfile
}

// NewGateway creates new gateway on top of existing
func (g *Base) NewGateway(ctx context.Context, state insolar.NetworkState) network.Gateway {
	inslogger.FromContext(ctx).Infof("NewGateway %s", state.String())
	switch state {
	case insolar.NoNetworkState:
		g.Self = newNoNetwork(g)
	case insolar.CompleteNetworkState:
		g.Self = newComplete(g)
	case insolar.JoinerBootstrap:
		g.Self = newJoinerBootstrap(g)
	case insolar.DiscoveryBootstrap:
		g.Self = newDiscoveryBootstrap(g)
	case insolar.WaitConsensus:
		g.Self = newWaitConsensus(g)
	case insolar.WaitMinRoles:
		g.Self = newWaitMinRoles(g)
	default:
		panic("Try to switch network to unknown state. Memory of process is inconsistent.")
	}
	return g.Self
}

func (g *Base) Init(ctx context.Context) error {
	g.HostNetwork.RegisterRequestHandler(types.Authorize, g.HandleNodeAuthorizeRequest) // validate cert
	g.HostNetwork.RegisterRequestHandler(types.Bootstrap, g.HandleNodeBootstrapRequest) // provide joiner claim
	g.HostNetwork.RegisterRequestHandler(types.UpdateSchedule, g.HandleUpdateSchedule)
	g.HostNetwork.RegisterRequestHandler(types.Reconnect, g.HandleReconnect)
	g.HostNetwork.RegisterRequestHandler(types.Ping, func(ctx context.Context, req network.ReceivedPacket) (network.Packet, error) {
		return g.HostNetwork.BuildResponse(ctx, req, &packet.Ping{}), nil
	})

	var err error

	// todo: use candidate profile
	//g.joinClaim, err = g.NodeKeeper.GetOriginJoinClaim()
	return err
}

func (g *Base) OnPulse(ctx context.Context, pu insolar.Pulse) error {
	return nil
}

func (g *Base) NeedLockMessageBus() bool {
	return true
}

// Auther casts us to Auther or obtain it in another way
func (g *Base) Auther() network.Auther {
	if ret, ok := g.Self.(network.Auther); ok {
		return ret
	}
	panic("Our network gateway suddenly is not an Auther")
}

// Bootstrapper casts us to Bootstrapper or obtain it in another way
func (g *Base) Bootstrapper() network.Bootstrapper {
	if ret, ok := g.Self.(network.Bootstrapper); ok {
		return ret
	}
	panic("Our network gateway suddenly is not an Bootstrapper")
}

// GetCert method returns node certificate by requesting sign from discovery nodes
func (g *Base) GetCert(ctx context.Context, ref *insolar.Reference) (insolar.Certificate, error) {
	return nil, errors.New("GetCert() in non active mode")
}

// ValidateCert validates node certificate
func (g *Base) ValidateCert(ctx context.Context, certificate insolar.AuthorizationCertificate) (bool, error) {
	return g.CertificateManager.VerifyAuthorizationCertificate(certificate)
}

// ============= Bootstrap =======

func (g *Base) ShouldIgnorePulse(context.Context, insolar.Pulse) bool {
	return false
}

func (g *Base) HandleNodeBootstrapRequest(ctx context.Context, request network.ReceivedPacket) (network.Packet, error) {
	if request.GetRequest() == nil || request.GetRequest().GetBootstrap() == nil {
		return nil, errors.Errorf("process bootstrap: got invalid protobuf request message: %s", request)
	}

	data := request.GetRequest().GetBootstrap()
	//candidate := data.CandidateProfile

	if network.CheckShortIDCollision(g.NodeKeeper.GetAccessor().GetActiveNodes(), insolar.ShortNodeID(data.CandidateProfile.ShortID)) {
		return g.HostNetwork.BuildResponse(ctx, request, &packet.BootstrapResponse{Code: packet.UpdateShortID}), nil
	}

	// shortID := network.GenerateUniqueShortID(g.NodeKeeper.GetAccessor().GetActiveNodes(), data.JoinClaim.GetNodeID())
	// } else {
	// 	shortID = data.JoinClaim.ShortNodeID
	// }

	// data.LastNodePulse

	lastPulse, err := g.PulseAccessor.Latest(ctx)
	if err != nil {
		lastPulse = *insolar.GenesisPulse
	}
	if lastPulse.PulseNumber > data.Pulse.PulseNumber {
		return g.HostNetwork.BuildResponse(ctx, request, &packet.BootstrapResponse{Code: packet.UpdateSchedule}), nil
	}

	err = bootstrap.ValidatePermit(data.Permit, g.CertificateManager.GetCertificate(), g.CryptographyService)
	if err != nil {
		inslogger.FromContext(ctx).Errorf("Rejected bootstrap request from node %s: %s", request.GetSender(), err.Error())
		return g.HostNetwork.BuildResponse(ctx, request, &packet.BootstrapResponse{Code: packet.Reject}), nil
	}

	//TODO: how to ignore claim if node already bootstrap to other??

	// TODO: check JoinClaim is from Discovery node
	//g.NodeKeeper.GetClaimQueue().Push(data.CandidateProfile)

	go func() {
		// TODO:
		//pulseStartTime := time.Unix(0, data.Pulse.PulseTimestamp)

		//pulseStartTime := time.Now()
		//g.PulseAppender.Append(ctx, lastPulse)
		//if err = g.PhaseManager.OnPulse(ctx, &lastPulse, pulseStartTime); err != nil {
		//	inslogger.FromContext(ctx).Error("Failed to pass consensus: ", err.Error())
		//}
		//if err = g.NodeKeeper.MoveSyncToActive(ctx, lastPulse.PulseNumber); err != nil {
		//	inslogger.FromContext(ctx).Error("Failed to MoveSyncToActive: ", err.Error())
		//}

		// fixme twice consensus call
		//lastPulse.PulseNumber += 1
		//if err := g.PhaseManager.OnPulse(ctx, &lastPulse, pulseStartTime); err != nil {
		//	inslogger.FromContext(ctx).Error("Failed to pass consensus: ", err.Error())
		//}
	}()

	// networkSize := uint32(len(g.NodeKeeper.GetAccessor().GetActiveNodes()))
	return g.HostNetwork.BuildResponse(ctx, request,
		&packet.BootstrapResponse{
			Code:  packet.Accepted,
			Pulse: *pulse.ToProto(&lastPulse),
			ETA:   uint32(lastPulse.PulseNumber) + 50, // TODO: calculate ETA
		}), nil
}

// validateTimestamp returns true if difference between timestamp ant current UTC < delta
func validateTimestamp(timestamp int64, delta time.Duration) bool {
	return time.Now().UTC().Sub(time.Unix(timestamp, 0)) < delta
}

func (g *Base) HandleNodeAuthorizeRequest(ctx context.Context, request network.ReceivedPacket) (network.Packet, error) {
	if !network.OriginIsDiscovery(g.CertificateManager.GetCertificate()) {
		return nil, errors.New("Only discovery nodes could authorize other nodes. I am not a discovery node.")
	}

	if request.GetRequest() == nil || request.GetRequest().GetAuthorize() == nil {
		return nil, errors.Errorf("process authorize: got invalid protobuf request message: %s", request)
	}
	data := request.GetRequest().GetAuthorize().AuthorizeData

	// TODO: move time.Minute to config
	if !validateTimestamp(data.Timestamp, time.Minute) {
		return g.HostNetwork.BuildResponse(ctx, request, &packet.AuthorizeResponse{
			Code:      packet.WrongTimestamp,
			Timestamp: time.Now().UTC().Unix(),
		}), nil
	}

	cert, err := certificate.Deserialize(data.Certificate, platformpolicy.NewKeyProcessor())
	if err != nil {
		return g.HostNetwork.BuildResponse(ctx, request, &packet.AuthorizeResponse{Code: packet.WrongMandate, Error: err.Error()}), nil
	}

	valid, err := g.Gatewayer.Gateway().Auther().ValidateCert(ctx, cert)
	if !valid {
		if err == nil {
			err = errors.New("Certificate validation failed")
		}

		// FIXME
		//return g.HostNetwork.BuildResponse(ctx, request, &packet.AuthorizeResponse{Code: packet.WrongMandate, Error: err.Error()}), nil
	}

	// TODO: get random reconnectHost
	// nodes := g.NodeKeeper.GetAccessor().GetActiveNodes()
	o := g.NodeKeeper.GetOrigin()

	// workaround bootstrap to the origin node
	reconnectHost, err := host.NewHostNS(o.Address(), o.ID(), o.ShortID())
	if err != nil {
		err = errors.Wrap(err, "Failed to get reconnectHost")
		inslogger.FromContext(ctx).Error(err.Error())
		return nil, err
	}

	pubKey, err := g.KeyProcessor.ExportPublicKeyPEM(o.PublicKey())
	if err != nil {
		err = errors.Wrap(err, "Failed to export public key")
		inslogger.FromContext(ctx).Error(err.Error())
		return nil, err
	}

	permit, err := bootstrap.CreatePermit(g.NodeKeeper.GetOrigin().ID(),
		reconnectHost,
		pubKey,
		g.CryptographyService,
	)
	if err != nil {
		return nil, err
	}

	p, err := g.PulseAccessor.Latest(ctx)
	if err != nil {
		inslogger.FromContext(ctx).Warn("Ephemeral pulse")
		p = *insolar.EphemeralPulse
	}

	discoveryCount := len(network.FindDiscoveriesInNodeList(g.NodeKeeper.GetAccessor().GetActiveNodes(), g.CertificateManager.GetCertificate()))
	return g.HostNetwork.BuildResponse(ctx, request, &packet.AuthorizeResponse{
		Code:           packet.Success,
		Timestamp:      time.Now().UTC().Unix(),
		Permit:         permit,
		DiscoveryCount: uint32(discoveryCount),
		Pulse:          pulse.ToProto(&p),
		//NetworkState:   uint32(g.Gatewayer.Gateway().GetState()),
	}), nil
}

func (g *Base) HandleUpdateSchedule(ctx context.Context, request network.ReceivedPacket) (network.Packet, error) {
	storage.NewSnapshotStorage()
	// TODO:
	return g.HostNetwork.BuildResponse(ctx, request, &packet.UpdateScheduleResponse{}), nil
}

func (g *Base) HandleReconnect(ctx context.Context, request network.ReceivedPacket) (network.Packet, error) {
	if request.GetRequest() == nil || request.GetRequest().GetReconnect() == nil {
		return nil, errors.Errorf("process reconnect: got invalid protobuf request message: %s", request)
	}

	// check permit, if permit from Discovery node
	// request.GetRequest().GetReconnect().Permit

	// TODO:
	return g.HostNetwork.BuildResponse(ctx, request, &packet.ReconnectResponse{}), nil
}

func (g *Base) OnConsensusFinished(p insolar.PulseNumber) {
	log.Infof("================== OnConsensusFinished for pulse %d", p)
}
