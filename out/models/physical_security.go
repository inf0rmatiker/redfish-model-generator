/* -----------------------------------------------------------------
* physical_security.go -
*
* DMTF Redfish PhysicalSecurity resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The state of the physical security sensor.
type PhysicalSecurity struct {
	// This indicates the known state of the physical security sensor, such as if it is hardware intrusion detected.
	IntrusionSensor IntrusionSensor `json:"IntrusionSensor,omitempty"`

	// A numerical identifier to represent the physical security sensor.
	IntrusionSensorNumber int `json:"IntrusionSensorNumber,omitempty"`

	// The method that restores this physical security sensor to the normal state.
	IntrusionSensorReArm IntrusionSensorReArm `json:"IntrusionSensorReArm,omitempty"`

}
