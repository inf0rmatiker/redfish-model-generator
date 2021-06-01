/* -----------------------------------------------------------------
* enum_read_cache_policy_type.go -
*
* SNIA Swordfish ReadCachePolicyType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReadCachePolicyType string

const (
	// A caching technique in which the controller pre-fetches data anticipating future read requests.
	ReadCachePolicyType_READ_AHEAD ReadCachePolicyType = "ReadAhead"

	// A caching technique in which the controller dynamically determines whether to pre-fetch data anticipating future read requests, based on previous cache hit ratio.
	ReadCachePolicyType_ADAPTIVE_READ_AHEAD ReadCachePolicyType = "AdaptiveReadAhead"

	// The read cache is disabled.
	ReadCachePolicyType_OFF ReadCachePolicyType = "Off"

)
