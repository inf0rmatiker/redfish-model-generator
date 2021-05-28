/* -----------------------------------------------------------------
 * enum_status_indicator.go -
 *
 * DMTF Redfish StatusIndicator enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type StatusIndicator string

const (
	// The drive is OK.
	StatusIndicator_OK StatusIndicator = "OK"

	// The drive has failed.
	StatusIndicator_FAIL StatusIndicator = "Fail"

	// The drive is being rebuilt.
	StatusIndicator_REBUILD StatusIndicator = "Rebuild"

	// The drive still works but is predicted to fail soon.
	StatusIndicator_PREDICTIVE_FAILURE_ANALYSIS StatusIndicator = "PredictiveFailureAnalysis"

	// The drive has been marked to automatically rebuild and replace a failed drive.
	StatusIndicator_HOTSPARE StatusIndicator = "Hotspare"

	// The array to which this drive belongs has been degraded.
	StatusIndicator_IN_A_CRITICAL_ARRAY StatusIndicator = "InACriticalArray"

	// The array to which this drive belongs has failed.
	StatusIndicator_IN_A_FAILED_ARRAY StatusIndicator = "InAFailedArray"

)
