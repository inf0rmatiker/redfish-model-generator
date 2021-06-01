/* -----------------------------------------------------------------
* enum_nv_me_controller_type.go -
*
* DMTF Redfish NVMeControllerType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type NVMeControllerType string

const (
	// The NVMe controller is an admin controller.
	NVMeControllerType_ADMIN NVMeControllerType = "Admin"

	// The NVMe controller is a discovery controller.
	NVMeControllerType_DISCOVERY NVMeControllerType = "Discovery"

	// The NVMe controller is an IO controller.
	NVMeControllerType_IO NVMeControllerType = "IO"

)
