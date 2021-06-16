/*
 * Connectivity Service Configuration
 *
 * APIs to configure connectivity service in Aether Network
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package configapi

type SliceQos struct {

	// uplink data rate in bps
	Uplink int32 `json:"uplink,omitempty"`

	// downlink data rate in bps
	Downlink int32 `json:"downlink,omitempty"`

	// QCI/QFI for the traffic
	TrafficClass string `json:"traffic-class,omitempty"`
}
