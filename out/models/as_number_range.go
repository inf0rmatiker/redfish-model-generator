/* -----------------------------------------------------------------
* as_number_range.go -
*
* DMTF Redfish ASNumberRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Autonomous System (AS) number range.
type ASNumberRange struct {
	// Lower Autonomous System (AS) number.
	Lower int `json:"Lower,omitempty"`

	// Upper Autonomous System (AS) number.
	Upper int `json:"Upper,omitempty"`

}
