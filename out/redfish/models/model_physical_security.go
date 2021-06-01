/* -----------------------------------------------------------------
* physical_security.go -
*
* DMTF Redfish PhysicalSecurity resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PhysicalSecurity - The state of the physical security sensor.
// Modeled after DMTF Redfish schema PhysicalSecurity
type PhysicalSecurity struct {
	// A numerical identifier to represent the physical security sensor.
	IntrusionSensorNumber float64 `json:"IntrusionSensorNumber,omitempty"`

	// This indicates the known state of the physical security sensor, such as if it is hardware intrusion detected.
	IntrusionSensor IntrusionSensor `json:"IntrusionSensor,omitempty"`

	// This indicates how the Normal state to be restored.
	IntrusionSensorReArm IntrusionSensorReArm `json:"IntrusionSensorReArm,omitempty"`

}
