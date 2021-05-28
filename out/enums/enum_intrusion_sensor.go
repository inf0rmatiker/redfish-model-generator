/* -----------------------------------------------------------------
 * enum_intrusion_sensor.go -
 *
 * DMTF Redfish IntrusionSensor enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type IntrusionSensor string

const (
	// No abnormal physical security condition is detected at this time.
	IntrusionSensor_NORMAL IntrusionSensor = "Normal"

	// A door, lock, or other mechanism protecting the internal system hardware from being accessed is detected to be in an insecure state.
	IntrusionSensor_HARDWARE_INTRUSION IntrusionSensor = "HardwareIntrusion"

	// Physical tampering of the monitored entity is detected.
	IntrusionSensor_TAMPERING_DETECTED IntrusionSensor = "TamperingDetected"

)
