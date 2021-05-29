/* -----------------------------------------------------------------
* measurement_block.go -
*
* DMTF Redfish MeasurementBlock resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The DSP0274-defined measurement block information.
type MeasurementBlock struct {
	// The hexadecimal string representation of the numeric value of the DSP0274-defined Measurement field  of the measurement block.
	Measurement string `json:"Measurement,omitempty"`

	// The DSP0274-defined MeasurementSize field of the measurement block.
	MeasurementSize int `json:"MeasurementSize,omitempty"`

	// The DSP0274-defined MeasurementSpecification field of the measurement block.
	MeasurementSpecification int `json:"MeasurementSpecification,omitempty"`

}
