/* -----------------------------------------------------------------
* power_sensors.go -
*
* DMTF Redfish PowerSensors resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerSensors - This property contains the power sensors.
// Modeled after DMTF Redfish schema PowerSensors
type PowerSensors struct {
	// The Line 1 to Line 2 power reading for this circuit.
	Line1ToLine2 map[string]interface{} `json:"Line1ToLine2,omitempty"`

	// The Line 1 to Neutral power reading for this circuit.
	Line1ToNeutral map[string]interface{} `json:"Line1ToNeutral,omitempty"`

	// The Line 2 to Line 3 power reading for this circuit.
	Line2ToLine3 map[string]interface{} `json:"Line2ToLine3,omitempty"`

	// The Line 2 to Neutral power reading for this circuit.
	Line2ToNeutral map[string]interface{} `json:"Line2ToNeutral,omitempty"`

	// The Line 3 to Line 1 power reading for this circuit.
	Line3ToLine1 map[string]interface{} `json:"Line3ToLine1,omitempty"`

	// The Line 3 to Neutral power reading for this circuit.
	Line3ToNeutral map[string]interface{} `json:"Line3ToNeutral,omitempty"`

}
