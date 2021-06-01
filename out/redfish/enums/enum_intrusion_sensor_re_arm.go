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
	// This sensor would be restored to the Normal state by a manual re-arm.
	IntrusionSensorReArm_MANUAL IntrusionSensorReArm = "Manual"

	// This sensor would be restored to the Normal state automatically as no abnormal physical security conditions are detected.
	IntrusionSensorReArm_AUTOMATIC IntrusionSensorReArm = "Automatic"

)
