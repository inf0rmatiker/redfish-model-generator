/* -----------------------------------------------------------------
* temperature_summary.go -
*
* DMTF Redfish TemperatureSummary resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// TemperatureSummary - The temperature sensors for a subsystem.
// Modeled after DMTF Redfish schema TemperatureSummary
type TemperatureSummary struct {
	// The ambient temperature of this subsystem.
	Ambient map[string]interface{} `json:"Ambient,omitempty"`

	// The exhaust temperature of this subsystem.
	Exhaust map[string]interface{} `json:"Exhaust,omitempty"`

	// The intake temperature of this subsystem.
	Intake map[string]interface{} `json:"Intake,omitempty"`

	// The internal temperature of this subsystem.
	Internal map[string]interface{} `json:"Internal,omitempty"`

}
