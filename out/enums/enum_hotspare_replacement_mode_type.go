/* -----------------------------------------------------------------
 * enum_hotspare_replacement_mode_type.go -
 *
 * DMTF Redfish HotspareReplacementModeType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type HotspareReplacementModeType string

const (
	// The hot spare drive that is commissioned due to a drive failure reverts to a hot spare after the failed drive is replaced and rebuilt.
	HotspareReplacementModeType_REVERTIBLE HotspareReplacementModeType = "Revertible"

	// The hot spare drive that is commissioned due to a drive failure remains as a data drive and does not revert to a hot spare if the failed drive is replaced.
	HotspareReplacementModeType_NON_REVERTIBLE HotspareReplacementModeType = "NonRevertible"

)
