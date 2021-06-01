/* -----------------------------------------------------------------
* enum_over_write_policy.go -
*
* DMTF Redfish OverWritePolicy enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type OverWritePolicy string

const (
	// Completed tasks are not automatically overwritten.
	OverWritePolicy_MANUAL OverWritePolicy = "Manual"

	// Oldest completed tasks are overwritten.
	OverWritePolicy_OLDEST OverWritePolicy = "Oldest"

)
