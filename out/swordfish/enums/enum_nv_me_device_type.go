/* -----------------------------------------------------------------
* enum_nv_me_device_type.go -
*
* SNIA Swordfish NVMeDeviceType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type NVMeDeviceType string

const (
	// Specifies an  device type of Drive, indicating a NVMe device that presents as an NVMe SSD device.
	NVMeDeviceType_DRIVE NVMeDeviceType = "Drive"

	// Specifies an  device type of JBOF, indicating a NVMe device that presents as an NVMe smart enclosure for NVMe devices, typically NVMe Drives.
	NVMeDeviceType_JBOF NVMeDeviceType = "JBOF"

	// Specifies an  NVMe device type of FabricAttachArray, indicating a NVMe device that presents an NVMe front-end that abstracts the back end storage, typically with multiple options for availability and protection.
	NVMeDeviceType_FABRIC_ATTACH_ARRAY NVMeDeviceType = "FabricAttachArray"

)
