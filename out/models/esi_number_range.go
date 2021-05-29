/* -----------------------------------------------------------------
* esi_number_range.go -
*
* DMTF Redfish ESINumberRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Ethernet Segment Identifier (ESI) number range for an Ethernet fabric.
type ESINumberRange struct {
	// Lower Ethernet Segment Identifier (ESI) number.
	Lower int `json:"Lower,omitempty"`

	// Upper Ethernet Segment Identifier (ESI) number.
	Upper int `json:"Upper,omitempty"`

}
