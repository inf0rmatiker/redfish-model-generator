/* -----------------------------------------------------------------
* input_range.go -
*
* DMTF Redfish InputRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes an input range for a power supply.
type InputRange struct {
	// The maximum capacity of this power supply when operating in this input range.
	CapacityWatts float64 `json:"CapacityWatts,omitempty"`

	// The input voltage range.
	NominalVoltageType NominalVoltageType `json:"NominalVoltageType,omitempty"`

}
