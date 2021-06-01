/* -----------------------------------------------------------------
* common_bgp_properties.go -
*
* DMTF Redfish CommonBGPProperties resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CommonBGPProperties - Common BGP properties.
// Modeled after DMTF Redfish schema CommonBGPProperties
type CommonBGPProperties struct {
	// Autonomous System (AS) number range.
	ASNumberRange ASNumberRange `json:"ASNumberRange,omitempty"`

	// Border Gateway Protocol (BGP) neighbor related properties.
	BGPNeighbor BGPNeighbor `json:"BGPNeighbor,omitempty"`

	// Border Gateway Protocol (BGP) route related properties.
	BGPRoute BGPRoute `json:"BGPRoute,omitempty"`

	// Graceful restart related properties.
	GracefulRestart GracefulRestart `json:"GracefulRestart,omitempty"`

	// Multiple path related properties.
	MultiplePaths MultiplePaths `json:"MultiplePaths,omitempty"`

	// This property shall indicate whether community attributes are sent.
	SendCommunityEnabled bool `json:"SendCommunityEnabled,omitempty"`

}
