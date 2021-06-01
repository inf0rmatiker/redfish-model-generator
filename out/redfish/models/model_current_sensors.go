/* -----------------------------------------------------------------
* current_sensors.go -
*
* DMTF Redfish CurrentSensors resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CurrentSensors - The current sensors for this circuit.
// Modeled after DMTF Redfish schema CurrentSensors
type CurrentSensors struct {
	// Line 1 current sensor.
	Line1 map[string]interface{} `json:"Line1,omitempty"`

	// Line 2 current sensor.
	Line2 map[string]interface{} `json:"Line2,omitempty"`

	// Line 3 current sensor.
	Line3 map[string]interface{} `json:"Line3,omitempty"`

	// Neutral line current sensor.
	Neutral map[string]interface{} `json:"Neutral,omitempty"`

}
