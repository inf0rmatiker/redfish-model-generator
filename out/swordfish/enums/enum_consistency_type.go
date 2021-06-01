/* -----------------------------------------------------------------
* enum_consistency_type.go -
*
* SNIA Swordfish ConsistencyType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ConsistencyType string

const (
	// Sequentially consistent.
	ConsistencyType_SEQUENTIALLY_CONSISTENT ConsistencyType = "SequentiallyConsistent"

)
