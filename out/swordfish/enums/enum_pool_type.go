/* -----------------------------------------------------------------
* enum_pool_type.go -
*
* SNIA Swordfish PoolType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PoolType string

const (
	// This pool is of type block.
	PoolType_BLOCK PoolType = "Block"

	// This pool is of type file.
	PoolType_FILE PoolType = "File"

	// This pool is of type object.
	PoolType_OBJECT PoolType = "Object"

	// This pool is of type pool, indicating a hierarchy.
	PoolType_POOL PoolType = "Pool"

)
