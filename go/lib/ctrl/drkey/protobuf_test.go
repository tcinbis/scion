// Copyright 2020 ETH Zurich
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

package drkey_test

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/scionproto/scion/go/lib/addr"
	ctrl "github.com/scionproto/scion/go/lib/ctrl/drkey"
	"github.com/scionproto/scion/go/lib/drkey"
	"github.com/scionproto/scion/go/lib/util"
	"github.com/scionproto/scion/go/lib/xtest"
	dkpb "github.com/scionproto/scion/go/pkg/proto/drkey"
)

func TestLvl1reqToProtoRequest(t *testing.T) {
	now := time.Now().UTC()

	valTime, err := ptypes.TimestampProto(now)
	require.NoError(t, err)
	timestamp, err := ptypes.TimestampProto(now)
	require.NoError(t, err)

	pbReq := &dkpb.DRKeyLvl1Request{
		ValTime:   valTime,
		Timestamp: timestamp,
	}

	lvl1Req := ctrl.Lvl1Req{
		ValTime:   now,
		Timestamp: now,
	}

	req, err := ctrl.Lvl1reqToProtoRequest(lvl1Req)
	require.NoError(t, err)
	assert.Equal(t, pbReq, req)
}

func TestRequestToLvl1Req(t *testing.T) {
	now := time.Now().UTC()

	valTime, err := ptypes.TimestampProto(now)
	require.NoError(t, err)
	timestamp, err := ptypes.TimestampProto(now)
	require.NoError(t, err)

	req := &dkpb.DRKeyLvl1Request{
		ValTime:   valTime,
		Timestamp: timestamp,
	}

	lvl1Req, err := ctrl.RequestToLvl1Req(req)
	require.NoError(t, err)
	assert.Equal(t, now, lvl1Req.ValTime)
	assert.Equal(t, now, lvl1Req.Timestamp)
}

func TestKeyToLvl1Resp(t *testing.T) {
	epochBegin, err := ptypes.TimestampProto(util.SecsToTime(0))
	require.NoError(t, err)
	epochEnd, err := ptypes.TimestampProto(util.SecsToTime(1))
	require.NoError(t, err)
	dstIA := xtest.MustParseIA("1-ff00:0:110")
	srcIA := xtest.MustParseIA("1-ff00:0:111")
	k := xtest.MustParseHexString("c584cad32613547c64823c756651b6f5") // just a level 1 key

	lvl1Key := drkey.Lvl1Key{
		Lvl1Meta: drkey.Lvl1Meta{
			Epoch: drkey.NewEpoch(0, 1),
			SrcIA: srcIA,
			DstIA: dstIA,
		},
		Key: k,
	}

	targetResp := &dkpb.DRKeyLvl1Response{
		EpochBegin: epochBegin,
		EpochEnd:   epochEnd,
		Drkey:      k,
	}

	pbResp, err := ctrl.KeyToLvl1Resp(lvl1Key)
	targetResp.Timestamp = pbResp.Timestamp
	require.NoError(t, err)
	assert.Equal(t, targetResp, pbResp)

}

func TestGetLvl1KeyFromReply(t *testing.T) {
	epochBegin, err := ptypes.TimestampProto(util.SecsToTime(0))
	require.NoError(t, err)
	epochEnd, err := ptypes.TimestampProto(util.SecsToTime(1))
	require.NoError(t, err)
	dstIA := xtest.MustParseIA("1-ff00:0:110")
	srcIA := xtest.MustParseIA("1-ff00:0:111")
	k := xtest.MustParseHexString("c584cad32613547c64823c756651b6f5") // just a level 1 key

	resp := &dkpb.DRKeyLvl1Response{
		EpochBegin: epochBegin,
		EpochEnd:   epochEnd,
		Drkey:      k,
	}

	targetLvl1Key := drkey.Lvl1Key{
		Lvl1Meta: drkey.Lvl1Meta{
			Epoch: drkey.NewEpoch(0, 1),
			SrcIA: srcIA,
			DstIA: dstIA,
		},
		Key: k,
	}

	lvl1Key, err := ctrl.GetLvl1KeyFromReply(srcIA, dstIA, resp)
	require.NoError(t, err)
	assert.Equal(t, targetLvl1Key, lvl1Key)

}

func TestRequestToLvl2Req(t *testing.T) {
	now := time.Now().UTC()
	valTime, err := ptypes.TimestampProto(now)
	require.NoError(t, err)
	dstIA := xtest.MustParseIA("1-ff00:0:110")
	srcIA := xtest.MustParseIA("1-ff00:0:111")
	reqType := drkey.Host2Host
	hostType := addr.HostTypeSVC
	hostAddr := addr.SvcCS

	req := &dkpb.DRKeyLvl2Request{
		Protocol: "piskes",
		ReqType:  uint32(reqType),
		DstIa:    uint64(dstIA.IAInt()),
		SrcIa:    uint64(srcIA.IAInt()),
		ValTime:  valTime,
		SrcHost: &dkpb.DRKeyLvl2Request_DRKeyHost{
			Type: uint32(hostType),
			Host: hostAddr.Pack(),
		},
		DstHost: &dkpb.DRKeyLvl2Request_DRKeyHost{
			Type: uint32(hostType),
			Host: hostAddr.Pack(),
		},
	}

	targetLvl2Req := ctrl.Lvl2Req{
		Protocol: "piskes",
		ReqType:  uint32(reqType),
		ValTime:  now,
		SrcIA:    srcIA,
		DstIA:    dstIA,
		SrcHost: ctrl.Host{
			Type: hostType,
			Host: hostAddr.Pack(),
		},
		DstHost: ctrl.Host{
			Type: hostType,
			Host: hostAddr.Pack(),
		},
	}

	lvl2Req, err := ctrl.RequestToLvl2Req(req)
	require.NoError(t, err)
	assert.Equal(t, targetLvl2Req, lvl2Req)

}

func TestKeyToLvl2Resp(t *testing.T) {
	srcIA := xtest.MustParseIA("1-ff00:0:1")
	dstIA := xtest.MustParseIA("1-ff00:0:2")
	key := xtest.MustParseHexString("47bfbb7d94706dc9e79825e5a837b006")
	meta := drkey.Lvl2Meta{
		KeyType:  drkey.AS2AS,
		Protocol: "scmp",
		Epoch:    drkey.NewEpoch(0, 1),
		SrcIA:    srcIA,
		DstIA:    dstIA,
		SrcHost:  addr.HostNone{},
		DstHost:  addr.HostNone{},
	}
	lvl2Key := drkey.Lvl2Key{
		Lvl2Meta: meta,
		Key:      key,
	}

	epochBegin, err := ptypes.TimestampProto(lvl2Key.Epoch.NotBefore)
	require.NoError(t, err)
	epochEnd, err := ptypes.TimestampProto(lvl2Key.Epoch.NotAfter)
	require.NoError(t, err)

	targetResp := &dkpb.DRKeyLvl2Response{
		EpochBegin: epochBegin,
		EpochEnd:   epochEnd,
		Drkey:      []byte(lvl2Key.Key),
	}

	resp, err := ctrl.KeyToLvl2Resp(lvl2Key)
	require.NoError(t, err)
	targetResp.Timestamp = resp.Timestamp
	assert.Equal(t, targetResp, resp)
}
