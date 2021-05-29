/* -----------------------------------------------------------------
* temperature_summary.go -
*
* DMTF Redfish TemperatureSummary resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The temperature sensors for a subsystem.
type TemperatureSummary struct {
	// The ambient temperature of this subsystem.
	Ambient SensorExcerpt `json:"Ambient,omitempty"`

	// The exhaust temperature of this subsystem.
	Exhaust SensorExcerpt `json:"Exhaust,omitempty"`

	// The intake temperature of this subsystem.
	Intake SensorExcerpt `json:"Intake,omitempty"`

	// The internal temperature of this subsystem.
	Internal SensorExcerpt `json:"Internal,omitempty"`

}
