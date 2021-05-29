/* -----------------------------------------------------------------
* current_sensors.go -
*
* DMTF Redfish CurrentSensors resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The current sensors for this circuit.
type CurrentSensors struct {
	// Line 1 current sensor.
	Line1 SensorCurrentExcerpt `json:"Line1,omitempty"`

	// Line 2 current sensor.
	Line2 SensorCurrentExcerpt `json:"Line2,omitempty"`

	// Line 3 current sensor.
	Line3 SensorCurrentExcerpt `json:"Line3,omitempty"`

	// Neutral line current sensor.
	Neutral SensorCurrentExcerpt `json:"Neutral,omitempty"`

}
