/* -----------------------------------------------------------------
* enum_write_cache_policy_type.go -
*
* SNIA Swordfish WriteCachePolicyType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type WriteCachePolicyType string

const (
	// A caching technique in which the completion of a write request is not signaled until data is safely stored on non-volatile media.
	WriteCachePolicyType_WRITE_THROUGH WriteCachePolicyType = "WriteThrough"

	// A caching technique in which the completion of a write request is signaled as soon as the data is in cache, and actual writing to non-volatile media is guaranteed to occur at a later time.
	WriteCachePolicyType_PROTECTED_WRITE_BACK WriteCachePolicyType = "ProtectedWriteBack"

	// A caching technique in which the completion of a write request is signaled as soon as the data is in cache; actual writing to non-volatile media is not guaranteed to occur at a later time.
	WriteCachePolicyType_UNPROTECTED_WRITE_BACK WriteCachePolicyType = "UnprotectedWriteBack"

	// The write cache is disabled.
	WriteCachePolicyType_OFF WriteCachePolicyType = "Off"

)
