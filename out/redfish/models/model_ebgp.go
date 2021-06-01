/* -----------------------------------------------------------------
* ebgp.go -
*
* DMTF Redfish EBGP resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EBGP - External BGP (eBGP) related properties for an Ethernet fabric.
// Modeled after DMTF Redfish schema EBGP
type EBGP struct {
	// Autonomous System (AS) number range.
	ASNumberRange ASNumberRange `json:"ASNumberRange,omitempty"`

	// Allow duplicate Autonomous System (AS) path.
	AllowDuplicateASEnabled bool `json:"AllowDuplicateASEnabled,omitempty"`

	// Option to override an Autonomous System (AS) number with the AS number of the sending peer .
	AllowOverrideASEnabled bool `json:"AllowOverrideASEnabled,omitempty"`

	// Compare Multi Exit Discriminator (MED) status.
	AlwaysCompareMEDEnabled bool `json:"AlwaysCompareMEDEnabled,omitempty"`

	// Local preference value.
	BGPLocalPreference int `json:"BGPLocalPreference,omitempty"`

	// Border Gateway Protocol (BGP) neighbor related properties.
	BGPNeighbor BGPNeighbor `json:"BGPNeighbor,omitempty"`

	// Border Gateway Protocol (BGP) route related properties.
	BGPRoute BGPRoute `json:"BGPRoute,omitempty"`

	// BGP weight attribute.
	BGPWeight int `json:"BGPWeight,omitempty"`

	// Graceful restart related properties.
	GracefulRestart GracefulRestart `json:"GracefulRestart,omitempty"`

	// BGP Multi Exit Discriminator (MED) value.
	MED int `json:"MED,omitempty"`

	// External BGP (eBGP) multihop status.
	MultihopEnabled bool `json:"MultihopEnabled,omitempty"`

	// External BGP (eBGP) multihop Time to Live (TTL) value.
	MultihopTTL int `json:"MultihopTTL,omitempty"`

	// Multiple path related properties.
	MultiplePaths MultiplePaths `json:"MultiplePaths,omitempty"`

	// This property shall indicate whether community attributes are sent.
	SendCommunityEnabled bool `json:"SendCommunityEnabled,omitempty"`

}
