/* -----------------------------------------------------------------
* energy_sensors.go -
*
* DMTF Redfish EnergySensors resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The energy readings for this circuit.
type EnergySensors struct {
	// The Line 1 to Line 2 energy reading for this circuit.
	Line1ToLine2 SensorEnergykWhExcerpt `json:"Line1ToLine2,omitempty"`

	// The Line 1 to Neutral energy reading for this circuit.
	Line1ToNeutral SensorEnergykWhExcerpt `json:"Line1ToNeutral,omitempty"`

	// The Line 2 to Line 3 energy reading for this circuit.
	Line2ToLine3 SensorEnergykWhExcerpt `json:"Line2ToLine3,omitempty"`

	// The Line 2 to Neutral energy reading for this circuit.
	Line2ToNeutral SensorEnergykWhExcerpt `json:"Line2ToNeutral,omitempty"`

	// The Line 3 to Line 1 energy reading for this circuit.
	Line3ToLine1 SensorEnergykWhExcerpt `json:"Line3ToLine1,omitempty"`

	// The Line 3 to Neutral energy reading for this circuit.
	Line3ToNeutral SensorEnergykWhExcerpt `json:"Line3ToNeutral,omitempty"`

}
