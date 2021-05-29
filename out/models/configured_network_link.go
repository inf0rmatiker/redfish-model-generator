/* -----------------------------------------------------------------
* configured_network_link.go -
*
* DMTF Redfish ConfiguredNetworkLink resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A set of link settings that a port is configured to use for autonegotiation.
type ConfiguredNetworkLink struct {
	// The link speed per lane this port is configured to use for autonegotiation.
	ConfiguredLinkSpeedGbps float64 `json:"ConfiguredLinkSpeedGbps,omitempty"`

	// The link width this port is configured to use for autonegotiation in conjunction with the link speed.
	ConfiguredWidth int `json:"ConfiguredWidth,omitempty"`

}
