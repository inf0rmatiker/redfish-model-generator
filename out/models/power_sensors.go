/* -----------------------------------------------------------------
* power_sensors.go -
*
* DMTF Redfish PowerSensors resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This property contains the power sensors.
type PowerSensors struct {
	// The Line 1 to Line 2 power reading for this circuit.
	Line1ToLine2 SensorPowerExcerpt `json:"Line1ToLine2,omitempty"`

	// The Line 1 to Neutral power reading for this circuit.
	Line1ToNeutral SensorPowerExcerpt `json:"Line1ToNeutral,omitempty"`

	// The Line 2 to Line 3 power reading for this circuit.
	Line2ToLine3 SensorPowerExcerpt `json:"Line2ToLine3,omitempty"`

	// The Line 2 to Neutral power reading for this circuit.
	Line2ToNeutral SensorPowerExcerpt `json:"Line2ToNeutral,omitempty"`

	// The Line 3 to Line 1 power reading for this circuit.
	Line3ToLine1 SensorPowerExcerpt `json:"Line3ToLine1,omitempty"`

	// The Line 3 to Neutral power reading for this circuit.
	Line3ToNeutral SensorPowerExcerpt `json:"Line3ToNeutral,omitempty"`

}
