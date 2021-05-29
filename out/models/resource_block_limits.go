/* -----------------------------------------------------------------
* resource_block_limits.go -
*
* DMTF Redfish ResourceBlockLimits resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type specifies the allowable quantities of types of resource blocks for a composition request.
type ResourceBlockLimits struct {
	// The maximum number of resource blocks of type `Compute` allowed for the composition request.
	MaxCompute int `json:"MaxCompute,omitempty"`

	// The maximum number of resource blocks of type `ComputerSystem` allowed for the composition request.
	MaxComputerSystem int `json:"MaxComputerSystem,omitempty"`

	// The maximum number of resource blocks of type `Expansion` allowed for the composition request.
	MaxExpansion int `json:"MaxExpansion,omitempty"`

	// The maximum number of resource blocks of type `Memory` allowed for the composition request.
	MaxMemory int `json:"MaxMemory,omitempty"`

	// The maximum number of resource blocks of type `Network` allowed for the composition request.
	MaxNetwork int `json:"MaxNetwork,omitempty"`

	// The maximum number of resource blocks of type `Processor` allowed for the composition request.
	MaxProcessor int `json:"MaxProcessor,omitempty"`

	// The maximum number of resource blocks of type `Storage` allowed for the composition request.
	MaxStorage int `json:"MaxStorage,omitempty"`

	// The minimum number of resource blocks of type `Compute` required for the composition request.
	MinCompute int `json:"MinCompute,omitempty"`

	// The minimum number of resource blocks of type `ComputerSystem` required for the composition request.
	MinComputerSystem int `json:"MinComputerSystem,omitempty"`

	// The minimum number of resource blocks of type `Expansion` required for the composition request.
	MinExpansion int `json:"MinExpansion,omitempty"`

	// The minimum number of resource blocks of type `Memory` required for the composition request.
	MinMemory int `json:"MinMemory,omitempty"`

	// The minimum number of resource blocks of type `Network` required for the composition request.
	MinNetwork int `json:"MinNetwork,omitempty"`

	// The minimum number of resource blocks of type `Processor` required for the composition request.
	MinProcessor int `json:"MinProcessor,omitempty"`

	// The minimum number of resource blocks of type `Storage` required for the composition request.
	MinStorage int `json:"MinStorage,omitempty"`

}
