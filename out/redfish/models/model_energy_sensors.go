/* -----------------------------------------------------------------
* energy_sensors.go -
*
* DMTF Redfish EnergySensors resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EnergySensors - The energy readings for this circuit.
// Modeled after DMTF Redfish schema EnergySensors
type EnergySensors struct {
	// The Line 1 to Line 2 energy reading for this circuit.
	Line1ToLine2 map[string]interface{} `json:"Line1ToLine2,omitempty"`

	// The Line 1 to Neutral energy reading for this circuit.
	Line1ToNeutral map[string]interface{} `json:"Line1ToNeutral,omitempty"`

	// The Line 2 to Line 3 energy reading for this circuit.
	Line2ToLine3 map[string]interface{} `json:"Line2ToLine3,omitempty"`

	// The Line 2 to Neutral energy reading for this circuit.
	Line2ToNeutral map[string]interface{} `json:"Line2ToNeutral,omitempty"`

	// The Line 3 to Line 1 energy reading for this circuit.
	Line3ToLine1 map[string]interface{} `json:"Line3ToLine1,omitempty"`

	// The Line 3 to Neutral energy reading for this circuit.
	Line3ToNeutral map[string]interface{} `json:"Line3ToNeutral,omitempty"`

}
