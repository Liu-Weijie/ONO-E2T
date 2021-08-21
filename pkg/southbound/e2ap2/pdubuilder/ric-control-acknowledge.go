// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
)

func CreateRicControlAcknowledgeE2apPdu(ricReqID types.RicRequest, ranFuncID types.RanFunctionID,
	ricCtrlStatus e2apies.RiccontrolStatus) (*e2appdudescriptions.E2ApPdu, error) {

	ricRequestID := e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes29{
		Id:          int32(v2beta1.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicrequestId{
			RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ranFunctionID := e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes5{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RanfunctionId{
			Value: int32(ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricControlStatus := e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes24{
		Id:          int32(v2beta1.ProtocolIeIDRiccontrolStatus),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ricCtrlStatus,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicControl: &e2appdudescriptions.RicControl{
						SuccessfulOutcome: &e2appducontents.RiccontrolAcknowledge{
							ProtocolIes: &e2appducontents.RiccontrolAcknowledgeIes{
								E2ApProtocolIes29: &ricRequestID,  // RIC Requestor & RIC Instance ID
								E2ApProtocolIes5:  &ranFunctionID, // RAN function ID
								//E2ApProtocolIes20: &ricCallProcessID,  // RIC Call Process ID
								E2ApProtocolIes24: &ricControlStatus, // RIC Control Status
								//E2ApProtocolIes32: &ricControlOutcome, // RIC Control Outcome
							},
						},
						ProcedureCode: &e2ap_constants.IdRiccontrol{
							Value: int32(v2beta1.ProcedureCodeIDRICcontrol),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}