// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/codec"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/filter"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	//t.SkipNow()

	ctrl := gomock.NewController(t)
	conn := NewMockConn(ctrl)

	readCh := make(chan *e2appdudescriptions.E2ApPdu, 1)
	conn.EXPECT().Write(gomock.Any()).DoAndReturn(func(out []byte) (int, error) {
		_, err := codec.XER.Decode(out)
		if err != nil {
			return 0, nil
		}
		return len(out), nil
	}).AnyTimes()

	conn.EXPECT().Read(gomock.Any()).DoAndReturn(func(in []byte) (int, error) {
		res := <-readCh
		out, err := codec.XER.Encode(res)
		if err != nil {
			return 0, err
		}
		return copy(in, out), nil
	}).AnyTimes()

	meta := Metadata{
		ID:     "test",
		PlmnID: "onf",
	}
	channel := newChannel(context.Background(), conn, meta)
	assert.Equal(t, ID("test"), channel.ID())
	assert.Equal(t, ID("test"), channel.Metadata().ID)
	assert.Equal(t, PlmnID("onf"), channel.Metadata().PlmnID)

	req := newSubscribeRequest(1)
	err := channel.Send(req, codec.XER)
	assert.NoError(t, err)

	subCh := channel.Recv(context.TODO(), filter.RicSubscription(req.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value), codec.XER)

	readCh <- newSubscribeResponse(1)

	res := <-subCh
	assert.NotNil(t, res)

	req = newSubscribeRequest(2)
	go func() {
		time.Sleep(time.Second)
		readCh <- newSubscribeResponse(2)
	}()
	res, err = channel.SendRecv(context.TODO(), req, filter.RicSubscription(req.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value), codec.XER)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	ch := channel.Recv(context.TODO(), filter.RicIndication(req.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value), codec.XER)

	readCh <- newIndication(2)
	indication := <-ch
	assert.NotNil(t, indication)

	readCh <- newIndication(2)
	indication = <-ch
	assert.NotNil(t, indication)
}

func newSubscribeRequest(id types.RicRequestorID) *e2appdudescriptions.E2ApPdu {
	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsToBeSetup[100] = types.RicActionDef{
		RicActionID:         100,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W5MS,
		RicActionDefinition: []byte{0x11, 0x22},
	}

	ricActionsToBeSetup[200] = types.RicActionDef{
		RicActionID:         200,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_INSERT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_W10MS,
		RicActionDefinition: []byte{0x33, 0x44},
	}

	req, _ := pdubuilder.CreateRicSubscriptionRequestE2apPdu(
		types.RicRequest{RequestorID: id, InstanceID: 2},
		3, []byte{0x55, 0x66}, ricActionsToBeSetup)
	return req
}

func newSubscribeResponse(id int32) *e2appdudescriptions.E2ApPdu {
	res, _ := pdubuilder.CreateRicSubscriptionResponseE2apPdu(&types.RicRequest{RequestorID: types.RicRequestorID(id), InstanceID: 2}, 3, []*types.RicActionID{})
	return res
}

func newIndication(id int32) *e2appdudescriptions.E2ApPdu {
	ricInd := types.RicRequest{
		RequestorID: types.RicRequestorID(id),
		InstanceID:  2,
	}
	res, _ := pdubuilder.RicIndicationE2apPdu(ricInd, types.RanFunctionID(3), e2apies.RicactionType_RICACTION_TYPE_POLICY, types.RicIndicationSn(1), e2apies.RicindicationType_RICINDICATION_TYPE_REPORT,
		types.RicIndicationHeader("foo"), types.RicIndicationMessage("bar"), types.RicCallProcessID("baz"))
	return res
}
