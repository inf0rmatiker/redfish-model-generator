/* -----------------------------------------------------------------
* supported_link_capabilities.go -
*
* DMTF Redfish SupportedLinkCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The link capabilities of an associated port.
type SupportedLinkCapabilities struct {
	// An indication of whether the port is capable of autonegotiating speed.
	AutoSpeedNegotiation bool `json:"AutoSpeedNegotiation,omitempty"`

	// The set of link speed capabilities of this port.
	CapableLinkSpeedMbps []integer `json:"CapableLinkSpeedMbps,omitempty"`

	// The link network technology capabilities of this port.
	LinkNetworkTechnology LinkNetworkTechnology `json:"LinkNetworkTechnology,omitempty"`

	// The speed of the link in Mbit/s when this link network technology is active.
	LinkSpeedMbps int `json:"LinkSpeedMbps,omitempty"`

}
