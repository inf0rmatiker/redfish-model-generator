/* -----------------------------------------------------------------
* link_configuration.go -
*
* DMTF Redfish LinkConfiguration resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Properties of the link for which this port is configured.
type LinkConfiguration struct {
	// An indication of whether the port is capable of autonegotiating speed.
	AutoSpeedNegotiationCapable bool `json:"AutoSpeedNegotiationCapable,omitempty"`

	// Controls whether this port is configured to enable autonegotiating speed.
	AutoSpeedNegotiationEnabled bool `json:"AutoSpeedNegotiationEnabled,omitempty"`

	// The set of link speed capabilities of this port.
	CapableLinkSpeedGbps []number `json:"CapableLinkSpeedGbps,omitempty"`

	// The set of link speed and width pairs this port is configured to use for autonegotiation.
	ConfiguredNetworkLinks array `json:"ConfiguredNetworkLinks,omitempty"`

}
