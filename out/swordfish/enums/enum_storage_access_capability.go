/* -----------------------------------------------------------------
* enum_storage_access_capability.go -
*
* SNIA Swordfish StorageAccessCapability enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type StorageAccessCapability string

const (
	// Read.
	StorageAccessCapability_READ StorageAccessCapability = "Read"

	// Write Many.
	StorageAccessCapability_WRITE StorageAccessCapability = "Write"

	// WriteOnce.
	StorageAccessCapability_WRITE_ONCE StorageAccessCapability = "WriteOnce"

	// AppendOnly.
	StorageAccessCapability_APPEND StorageAccessCapability = "Append"

	// Streaming.
	StorageAccessCapability_STREAMING StorageAccessCapability = "Streaming"

	// Execute access is allowed by the file share.
	StorageAccessCapability_EXECUTE StorageAccessCapability = "Execute"

)
