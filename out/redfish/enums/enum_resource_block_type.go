/* -----------------------------------------------------------------
* enum_resource_block_type.go -
*
* DMTF Redfish ResourceBlockType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ResourceBlockType string

const (
	// This Resource Block contains both Processor and Memory resources in a manner that creates a compute complex.
	ResourceBlockType_COMPUTE ResourceBlockType = "Compute"

	// This Resource Block contains Processor resources.
	ResourceBlockType_PROCESSOR ResourceBlockType = "Processor"

	// This Resource Block contains Memory resources.
	ResourceBlockType_MEMORY ResourceBlockType = "Memory"

	// This Resource Block contains Network resources, such as Ethernet Interfaces.
	ResourceBlockType_NETWORK ResourceBlockType = "Network"

	// This Resource Block contains Storage resources, such as Storage and Simple Storage.
	ResourceBlockType_STORAGE ResourceBlockType = "Storage"

	// This Resource Block contains ComputerSystem resources.
	ResourceBlockType_COMPUTER_SYSTEM ResourceBlockType = "ComputerSystem"

	// This Resource Block is capable of changing over time based on its configuration.  Different types of devices within this Resource Block can be added and removed over time.
	ResourceBlockType_EXPANSION ResourceBlockType = "Expansion"

)
