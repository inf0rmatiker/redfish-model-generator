/* -----------------------------------------------------------------
* input_range.go -
*
* DMTF Redfish InputRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// The Input type (AC or DC).
	InputType InputType `json:"InputType,omitempty"`

	// The minimum line input voltage at which this power supply input range is effective.
	MinimumVoltage float64 `json:"MinimumVoltage,omitempty"`

	// The maximum line input voltage at which this power supply input range is effective.
	MaximumVoltage float64 `json:"MaximumVoltage,omitempty"`

	// The minimum line input frequency at which this power supply input range is effective.
	MinimumFrequencyHz float64 `json:"MinimumFrequencyHz,omitempty"`

	// The maximum line input frequency at which this power supply input range is effective.
	MaximumFrequencyHz float64 `json:"MaximumFrequencyHz,omitempty"`

	// The maximum capacity of this Power Supply when operating in this input range.
	OutputWattage float64 `json:"OutputWattage,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
