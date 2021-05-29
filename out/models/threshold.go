/* -----------------------------------------------------------------
* threshold.go -
*
* DMTF Redfish Threshold resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A threshold definition for a sensor.
type Threshold struct {
	// The direction of crossing that activates this threshold.
	Activation ThresholdActivation `json:"Activation,omitempty"`

	// The duration the sensor value must violate the threshold before the threshold is activated.
	DwellTime string `json:"DwellTime,omitempty"`

	// The threshold value.
	Reading float64 `json:"Reading,omitempty"`

}
