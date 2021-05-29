/* -----------------------------------------------------------------
* power_limit.go -
*
* DMTF Redfish PowerLimit resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The power limit status and configuration information for the chassis.
type PowerLimit struct {
	// The time required for the limiting process to reduce power consumption to below the limit.
	CorrectionInMs int `json:"CorrectionInMs,omitempty"`

	// The action that is taken if the power cannot be maintained below the LimitInWatts.
	LimitException PowerLimitException `json:"LimitException,omitempty"`

	// The power limit, in watts.  If `null`, power capping is disabled.
	LimitInWatts float64 `json:"LimitInWatts,omitempty"`

}
