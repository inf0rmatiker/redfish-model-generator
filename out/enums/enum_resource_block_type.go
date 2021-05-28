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
	// This resource block contains resources of type `Processor` and `Memory` in a manner that creates a compute complex.
	ResourceBlockType_COMPUTE ResourceBlockType = "Compute"

	// This resource block contains resources of type `Processor`.
	ResourceBlockType_PROCESSOR ResourceBlockType = "Processor"

	// This resource block contains resources of type `Memory`.
	ResourceBlockType_MEMORY ResourceBlockType = "Memory"

	// This resource block contains network resources, such as resource of type `EthernetInterface` and `NetworkInterface`.
	ResourceBlockType_NETWORK ResourceBlockType = "Network"

	// This resource block contains storage resources, such as resources of type `Storage` and `SimpleStorage`.
	ResourceBlockType_STORAGE ResourceBlockType = "Storage"

	// This resource block contains resources of type `ComputerSystem`.
	ResourceBlockType_COMPUTER_SYSTEM ResourceBlockType = "ComputerSystem"

	// This resource block is capable of changing over time based on its configuration.  Different types of devices within this resource block can be added and removed over time.
	ResourceBlockType_EXPANSION ResourceBlockType = "Expansion"

	// This resource block is capable of being consumed as a standalone component.  This resource block can represent things such as a software platform on one or more computer systems or an appliance that provides composable resources and other services, and can be managed independently of the Redfish service.
	ResourceBlockType_INDEPENDENT_RESOURCE ResourceBlockType = "IndependentResource"

)
