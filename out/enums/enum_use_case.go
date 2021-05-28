/* -----------------------------------------------------------------
 * enum_use_case.go -
 *
 * DMTF Redfish UseCase enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type UseCase string

const (
	// This capability describes a client creating a new computer system resource from a set of disaggregated hardware.
	UseCase_COMPUTER_SYSTEM_COMPOSITION UseCase = "ComputerSystemComposition"

	// This capability describes a client creating a new computer system resource from a set of constraints.
	UseCase_COMPUTER_SYSTEM_CONSTRAINED_COMPOSITION UseCase = "ComputerSystemConstrainedComposition"

	// This capability describes a client creating a new volume resource as part of an existing storage subsystem.
	UseCase_VOLUME_CREATION UseCase = "VolumeCreation"

	// This capability describes a client creating a new resource block from a set of other resource blocks.
	UseCase_RESOURCE_BLOCK_COMPOSITION UseCase = "ResourceBlockComposition"

	// This capability describes a client creating a new resource block from a set of constraints.
	UseCase_RESOURCE_BLOCK_CONSTRAINED_COMPOSITION UseCase = "ResourceBlockConstrainedComposition"

)
