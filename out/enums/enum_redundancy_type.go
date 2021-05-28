/* -----------------------------------------------------------------
 * enum_redundancy_type.go -
 *
 * DMTF Redfish RedundancyType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type RedundancyType string

const (
	// Failure of one unit automatically causes a standby or offline unit in the redundancy set to take over its functions.
	RedundancyType_FAILOVER RedundancyType = "Failover"

	// Multiple units are available and active such that normal operation will continue if one or more units fail.
	RedundancyType_N_PLUS_M RedundancyType = "NPlusM"

	// Multiple units contribute or share such that operation will continue, but at a reduced capacity, if one or more units fail.
	RedundancyType_SHARING RedundancyType = "Sharing"

	// One or more spare units are available to take over the function of a failed unit, but takeover is not automatic.
	RedundancyType_SPARING RedundancyType = "Sparing"

	// The subsystem is not configured in a redundancy mode, either due to configuration or the functionality has been disabled by the user.
	RedundancyType_NOT_REDUNDANT RedundancyType = "NotRedundant"

)
