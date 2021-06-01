/* -----------------------------------------------------------------
* enum_system_type.go -
*
* DMTF Redfish SystemType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SystemType string

const (
	// A computer system.
	SystemType_PHYSICAL SystemType = "Physical"

	// A virtual machine instance running on this system.
	SystemType_VIRTUAL SystemType = "Virtual"

	// An operating system instance.
	SystemType_OS SystemType = "OS"

	// A hardware-based partition of a computer system.
	SystemType_PHYSICALLY_PARTITIONED SystemType = "PhysicallyPartitioned"

	// A virtual or software-based partition of a computer system.
	SystemType_VIRTUALLY_PARTITIONED SystemType = "VirtuallyPartitioned"

	// A computer system constructed by binding resource blocks together.
	SystemType_COMPOSED SystemType = "Composed"

)
