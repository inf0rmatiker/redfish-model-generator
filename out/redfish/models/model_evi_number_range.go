/* -----------------------------------------------------------------
* evi_number_range.go -
*
* DMTF Redfish EVINumberRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EVINumberRange - The Ethernet Virtual Private Network (EVPN) Instance (EVI) number range for an Ethernet fabric.
// Modeled after DMTF Redfish schema EVINumberRange
type EVINumberRange struct {
	// Lower Ethernet Virtual Private Network (EVPN) Instance (EVI) number.
	Lower int `json:"Lower,omitempty"`

	// Upper Ethernet Virtual Private Network (EVPN) Instance (EVI) number.
	Upper int `json:"Upper,omitempty"`

}
