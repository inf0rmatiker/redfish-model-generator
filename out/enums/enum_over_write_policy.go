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
	// The overwrite policy is not known or is undefined.
	OverWritePolicy_UNKNOWN OverWritePolicy = "Unknown"

	// When full, new entries to the log overwrite earlier entries.
	OverWritePolicy_WRAPS_WHEN_FULL OverWritePolicy = "WrapsWhenFull"

	// When full, new entries to the log are discarded.
	OverWritePolicy_NEVER_OVER_WRITES OverWritePolicy = "NeverOverWrites"

)
