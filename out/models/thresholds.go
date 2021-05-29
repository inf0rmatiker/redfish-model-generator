/* -----------------------------------------------------------------
* thresholds.go -
*
* DMTF Redfish Thresholds resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The set of thresholds for a sensor.
type Thresholds struct {
	// The value at which the reading is below normal range and requires attention.
	LowerCritical Threshold `json:"LowerCritical,omitempty"`

	// The value at which the reading is below normal range.
	LowerWarning Threshold `json:"LowerWarning,omitempty"`

	// The value at which the reading is above normal range and requires attention.
	UpperCritical Threshold `json:"UpperCritical,omitempty"`

	// The value at which the reading is above normal range.
	UpperWarning Threshold `json:"UpperWarning,omitempty"`

}
