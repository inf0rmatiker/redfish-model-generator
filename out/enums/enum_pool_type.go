/* -----------------------------------------------------------------
 * enum_pool_type.go -
 *
 * DMTF Redfish PoolType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type PoolType string

const (
	// This resource block is in the free pool and is not contributing to any composed resources.
	PoolType_FREE PoolType = "Free"

	// This resource block is in the active pool and is contributing to at least one composed resource as a result of a composition request.
	PoolType_ACTIVE PoolType = "Active"

	// This resource block is not assigned to any pools.
	PoolType_UNASSIGNED PoolType = "Unassigned"

)
