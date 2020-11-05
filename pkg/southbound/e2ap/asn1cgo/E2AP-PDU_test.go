// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func Test_newE2setupResponseE2APpdu(t *testing.T) {
	e2SetupResponseE2APpdu, err := pdubuilder.CreateResponseE2apPdu("ONF", 0xABCDE, nil, nil)
	assert.NilError(t, err)

	e2SetupResponseE2APpduC, err := newE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	assert.Assert(t, e2SetupResponseE2APpduC != nil)

	xer, err := XerEncodeE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	t.Logf("XER of E2AP: %s\n", string(xer))

	per, err := PerEncodeE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	t.Logf("PER of E2AP: %v\n", string(per))

}