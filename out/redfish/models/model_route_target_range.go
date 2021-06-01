/* -----------------------------------------------------------------
* route_target_range.go -
*
* DMTF Redfish RouteTargetRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// RouteTargetRange - The Route Target (RT) number range for the fabric.
// Modeled after DMTF Redfish schema RouteTargetRange
type RouteTargetRange struct {
	// Lower Route Target (RT) number.
	Lower int `json:"Lower,omitempty"`

	// Upper Route Target (RT) number.
	Upper int `json:"Upper,omitempty"`

}
