// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

type RicRequestorID int32
type RicInstanceID int32

type RicRequest struct {
	RequestorID RicRequestorID
	InstanceID  RicInstanceID
}
