// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

/*
 * Connectivity Service Configuration
 *
 * APIs to configure connectivity service in Aether Network
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package configmodels

type SliceApplicationFilteringRules struct {

	// Rule name
	RuleName string `json:"rule-name,omitempty"`

	//priority
	Priority int32 `json:"priority,omitempty"`

	//action
	Action string `json:"action,omitempty"`

	// Application Desination IP or network
	Endpoint string `json:"endpoint,omitempty"`

	//protocol
	Protocol int32 `json:"protocol,omitempty"`

	// port range start
	StartPort int32 `json:"dest-port-start,omitempty"`

	// port range end
	EndPort int32 `json:"dest-port-end,omitempty"`

	AppMbrUplink int32 `json:"app-mbr-uplink,omitempty"`

	AppMbrDownlink int32 `json:"app-mbr-downlink,omitempty"`

	// data rate unit for uplink and downlink
	BitrateUnit string `json:"bitrate-unit,omitempty"`

	TrafficClass *TrafficClassInfo `json:"traffic-class,omitempty"`

	RuleTrigger string `json:"rule-trigger,omitempty"`
}
