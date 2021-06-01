/* -----------------------------------------------------------------
* enum_write_cache_state_type.go -
*
* DMTF Redfish WriteCacheStateType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type WriteCacheStateType string

const (
	// Indicates that the cache state type in use generally does not protect write requests on non-volatile media.
	WriteCacheStateType_UNPROTECTED WriteCacheStateType = "Unprotected"

	// Indicates that the cache state type in use generally protects write requests on non-volatile media.
	WriteCacheStateType_PROTECTED WriteCacheStateType = "Protected"

	// Indicates an issue with the cache state in which the cache space is diminished or disabled due to a failure or an outside influence such as a discharged battery.
	WriteCacheStateType_DEGRADED WriteCacheStateType = "Degraded"

)
