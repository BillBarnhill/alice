// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dkg

import (
	"errors"

	"github.com/getamis/alice/crypto/birkhoffinterpolation"
	"github.com/getamis/alice/crypto/ecpointgrouplaw"
	"github.com/getamis/alice/crypto/tss"
	"github.com/getamis/alice/crypto/tss/message/types"
	"github.com/getamis/sirius/log"
)

var (
	ErrInconsistentPubKey = errors.New("inconsistent public key")
)

type resultData struct {
	result *ecpointgrouplaw.ECPoint
}

type resultHandler struct {
	*verifyHandler
}

func newResultHandler(v *verifyHandler) *resultHandler {
	return &resultHandler{
		verifyHandler: v,
	}
}

func (p *resultHandler) MessageType() types.MessageType {
	return types.MessageType(Type_Result)
}

func (p *resultHandler) IsHandled(logger log.Logger, id string) bool {
	peer, ok := p.peers[id]
	if !ok {
		logger.Warn("Peer not found")
		return false
	}
	return peer.result != nil
}

func (p *resultHandler) HandleMessage(logger log.Logger, message types.Message) error {
	msg := getMessage(message)
	id := msg.GetId()
	peer, ok := p.peers[id]
	if !ok {
		logger.Warn("Peer not found")
		return tss.ErrPeerNotFound
	}

	siGProofMsg := msg.GetResult().SiGProofMsg
	r, err := siGProofMsg.V.ToPoint()
	if err != nil {
		logger.Warn("Failed to get point", "err", err)
		return err
	}
	err = siGProofMsg.Verify(ecpointgrouplaw.NewBase(p.publicKey.GetCurve()))
	if err != nil {
		logger.Warn("Failed to verify Schorr proof", "err", err)
		return err
	}
	peer.result = &resultData{
		result: r,
	}
	return peer.AddMessage(msg)
}

func (p *resultHandler) Finalize(logger log.Logger) (types.Handler, error) {
	bks := make(birkhoffinterpolation.BkParameters, p.peerNum+1)
	sgs := make([]*ecpointgrouplaw.ECPoint, p.peerNum+1)
	siG, err := p.siGProofMsg.V.ToPoint()
	if err != nil {
		logger.Warn("Failed to get point", "err", err)
		return nil, err
	}
	bks[0] = p.bk
	sgs[0] = siG
	i := 1
	for _, peer := range p.peers {
		bks[i] = peer.peer.bk
		sgs[i] = peer.result.result
		i++
	}
	scalars, err := bks.ComputeBkCoefficient(p.threshold, p.curve.Params().N)
	if err != nil {
		logger.Warn("Failed to compute", "err", err)
		return nil, err
	}
	gotPub, err := ecpointgrouplaw.ComputeLinearCombinationPoint(scalars, sgs)
	if err != nil {
		logger.Warn("Failed to calculate public key", "err", err)
		return nil, err
	}
	if !p.publicKey.Equal(gotPub) {
		logger.Warn("Inconsistent public key", "got", gotPub, "expected", p.publicKey)
		return nil, ErrInconsistentPubKey
	}
	return nil, nil
}
