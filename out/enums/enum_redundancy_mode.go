/* -----------------------------------------------------------------
 * enum_redundancy_mode.go -
 *
 * DMTF Redfish RedundancyMode enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type RedundancyMode string

const (
	// Failure of one unit automatically causes a standby or offline unit in the redundancy set to take over its functions.
	RedundancyMode_FAILOVER RedundancyMode = "Failover"

	// Multiple units are available and active such that normal operation will continue if one or more units fail.
	RedundancyMode_N+M RedundancyMode = "N+m"

	// Multiple units contribute or share such that operation will continue, but at a reduced capacity, if one or more units fail.
	RedundancyMode_SHARING RedundancyMode = "Sharing"

	// One or more spare units are available to take over the function of a failed unit, but takeover is not automatic.
	RedundancyMode_SPARING RedundancyMode = "Sparing"

	// The subsystem is not configured in a redundancy mode, either due to configuration or the functionality has been disabled by the user.
	RedundancyMode_NOT_REDUNDANT RedundancyMode = "NotRedundant"

)
