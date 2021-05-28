/* -----------------------------------------------------------------
 * enum_intrusion_sensor_re_arm.go -
 *
 * DMTF Redfish IntrusionSensorReArm enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type IntrusionSensorReArm string

const (
	// A manual re-arm of this sensor restores it to the normal state.
	IntrusionSensorReArm_MANUAL IntrusionSensorReArm = "Manual"

	// Because no abnormal physical security condition is detected, this sensor is automatically restored to the normal state.
	IntrusionSensorReArm_AUTOMATIC IntrusionSensorReArm = "Automatic"

)
