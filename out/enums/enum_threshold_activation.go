/* -----------------------------------------------------------------
 * enum_threshold_activation.go -
 *
 * DMTF Redfish ThresholdActivation enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ThresholdActivation string

const (
	// Value increases above the threshold.
	ThresholdActivation_INCREASING ThresholdActivation = "Increasing"

	// Value decreases below the threshold.
	ThresholdActivation_DECREASING ThresholdActivation = "Decreasing"

	// Value crosses the threshold in either direction.
	ThresholdActivation_EITHER ThresholdActivation = "Either"

)
