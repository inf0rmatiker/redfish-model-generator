/* -----------------------------------------------------------------
* threshold.go -
*
* DMTF Redfish Threshold resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Threshold - A threshold definition for a sensor.
// Modeled after DMTF Redfish schema Threshold
type Threshold struct {
	// The direction of crossing that activates this threshold.
	Activation ThresholdActivation `json:"Activation,omitempty"`

	// The time interval over which the sensor reading must have passed through this Threshold value before the threshold is considered to be violated.
	DwellTime string `json:"DwellTime,omitempty"`

	// The threshold value.
	Reading float64 `json:"Reading,omitempty"`

}
