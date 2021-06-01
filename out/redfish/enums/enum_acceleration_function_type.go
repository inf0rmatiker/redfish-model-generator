/* -----------------------------------------------------------------
* enum_acceleration_function_type.go -
*
* DMTF Redfish AccelerationFunctionType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AccelerationFunctionType string

const (
	// An encryption function.
	AccelerationFunctionType_ENCRYPTION AccelerationFunctionType = "Encryption"

	// A compression function.
	AccelerationFunctionType_COMPRESSION AccelerationFunctionType = "Compression"

	// A packet inspection function.
	AccelerationFunctionType_PACKET_INSPECTION AccelerationFunctionType = "PacketInspection"

	// A packet switch function.
	AccelerationFunctionType_PACKET_SWITCH AccelerationFunctionType = "PacketSwitch"

	// A scheduler function.
	AccelerationFunctionType_SCHEDULER AccelerationFunctionType = "Scheduler"

	// An audio processing function.
	AccelerationFunctionType_AUDIO_PROCESSING AccelerationFunctionType = "AudioProcessing"

	// A video processing function.
	AccelerationFunctionType_VIDEO_PROCESSING AccelerationFunctionType = "VideoProcessing"

	// An OEM-defined acceleration function.
	AccelerationFunctionType_OEM AccelerationFunctionType = "OEM"

)
