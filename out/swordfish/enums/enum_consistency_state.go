/* -----------------------------------------------------------------
* enum_consistency_state.go -
*
* SNIA Swordfish ConsistencyState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ConsistencyState string

const (
	// Consistent.
	ConsistencyState_CONSISTENT ConsistencyState = "Consistent"

	// Not consistent.
	ConsistencyState_INCONSISTENT ConsistencyState = "Inconsistent"

)
