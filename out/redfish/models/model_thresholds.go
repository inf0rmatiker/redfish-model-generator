/* -----------------------------------------------------------------
* thresholds.go -
*
* DMTF Redfish Thresholds resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Thresholds - The set of thresholds defined for a sensor.
// Modeled after DMTF Redfish schema Thresholds
type Thresholds struct {
	// Below normal range and requires attention.
	LowerCritical Threshold `json:"LowerCritical,omitempty"`

	// Below normal range.
	LowerWarning Threshold `json:"LowerWarning,omitempty"`

	// Above normal range and requires attention.
	UpperCritical Threshold `json:"UpperCritical,omitempty"`

	// Above normal range.
	UpperWarning Threshold `json:"UpperWarning,omitempty"`

}
