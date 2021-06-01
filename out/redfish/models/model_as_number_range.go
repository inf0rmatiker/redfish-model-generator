/* -----------------------------------------------------------------
* as_number_range.go -
*
* DMTF Redfish ASNumberRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ASNumberRange - Autonomous System (AS) number range.
// Modeled after DMTF Redfish schema ASNumberRange
type ASNumberRange struct {
	// Lower Autonomous System (AS) number.
	Lower int `json:"Lower,omitempty"`

	// Upper Autonomous System (AS) number.
	Upper int `json:"Upper,omitempty"`

}
