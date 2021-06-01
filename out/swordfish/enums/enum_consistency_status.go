/* -----------------------------------------------------------------
* enum_consistency_status.go -
*
* SNIA Swordfish ConsistencyStatus enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ConsistencyStatus string

const (
	// Consistent.
	ConsistencyStatus_CONSISTENT ConsistencyStatus = "Consistent"

	// Becoming consistent.
	ConsistencyStatus_IN_PROGRESS ConsistencyStatus = "InProgress"

	// Consistency disabled.
	ConsistencyStatus_DISABLED ConsistencyStatus = "Disabled"

	// Consistency error.
	ConsistencyStatus_IN_ERROR ConsistencyStatus = "InError"

)
