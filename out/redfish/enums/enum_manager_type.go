/* -----------------------------------------------------------------
* enum_manager_type.go -
*
* DMTF Redfish ManagerType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ManagerType string

const (
	// A controller used primarily to monitor or manage the operation of a device or system.
	ManagerType_MANAGEMENT_CONTROLLER ManagerType = "ManagementController"

	// A controller which provides management functions for a chassis or group of devices or systems.
	ManagerType_ENCLOSURE_MANAGER ManagerType = "EnclosureManager"

	// A controller which provides management functions for a single computer system.
	ManagerType_BMC ManagerType = "BMC"

	// A controller which provides management functions for a whole or part of a rack.
	ManagerType_RACK_MANAGER ManagerType = "RackManager"

	// A controller which provides management functions for a particular subsystem or group of devices.
	ManagerType_AUXILIARY_CONTROLLER ManagerType = "AuxiliaryController"

	// A software-based service which provides management functions.
	ManagerType_SERVICE ManagerType = "Service"

)