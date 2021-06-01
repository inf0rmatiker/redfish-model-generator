/* -----------------------------------------------------------------
* enum_device_type.go -
*
* DMTF Redfish DeviceType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type DeviceType string

const (
	// A single-function PCIe device.
	DeviceType_SINGLE_FUNCTION DeviceType = "SingleFunction"

	// A multi-function PCIe device.
	DeviceType_MULTI_FUNCTION DeviceType = "MultiFunction"

	// A PCIe device which is not currently physically present, but is being simulated by the PCIe infrastructure.
	DeviceType_SIMULATED DeviceType = "Simulated"

)
