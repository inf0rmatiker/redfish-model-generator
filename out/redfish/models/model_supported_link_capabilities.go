/* -----------------------------------------------------------------
* supported_link_capabilities.go -
*
* DMTF Redfish SupportedLinkCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SupportedLinkCapabilities - The self-described link capabilities of an assocaited port.
// Modeled after DMTF Redfish schema SupportedLinkCapabilities
type SupportedLinkCapabilities struct {
	// The self-described link network technology capabilities of this port.
	LinkNetworkTechnology LinkNetworkTechnology `json:"LinkNetworkTechnology,omitempty"`

	// The speed of the link in Mbps when this link network technology is active.
	LinkSpeedMbps float64 `json:"LinkSpeedMbps,omitempty"`

}
