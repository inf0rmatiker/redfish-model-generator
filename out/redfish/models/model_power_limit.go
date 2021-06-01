/* -----------------------------------------------------------------
* power_limit.go -
*
* DMTF Redfish PowerLimit resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerLimit - This object contains power limit status and configuration information for the chassis.
// Modeled after DMTF Redfish schema PowerLimit
type PowerLimit struct {
	// The Power limit in watts. Set to null to disable power capping.
	LimitInWatts float64 `json:"LimitInWatts,omitempty"`

	// The action that is taken if the power cannot be maintained below the LimitInWatts.
	LimitException PowerLimitException `json:"LimitException,omitempty"`

	// The time required for the limiting process to reduce power consumption to below the limit.
	CorrectionInMs float64 `json:"CorrectionInMs,omitempty"`

}
