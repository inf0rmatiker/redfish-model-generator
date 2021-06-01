/* -----------------------------------------------------------------
* output_rail.go -
*
* DMTF Redfish OutputRail resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// OutputRail - This type describes an output power rail for a power supply.
// Modeled after DMTF Redfish schema OutputRail
type OutputRail struct {
	// The nominal voltage of this output power rail.
	NominalVoltage float64 `json:"NominalVoltage,omitempty"`

	// The area or device to which this power rail applies.
	PhysicalContext map[string]interface{} `json:"PhysicalContext,omitempty"`

}
