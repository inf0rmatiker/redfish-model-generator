/* -----------------------------------------------------------------
* output_rail.go -
*
* DMTF Redfish OutputRail resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes an output power rail for a power supply.
type OutputRail struct {
	// The nominal voltage of this output power rail.
	NominalVoltage float64 `json:"NominalVoltage,omitempty"`

	// The area or device to which this power rail applies.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

}
