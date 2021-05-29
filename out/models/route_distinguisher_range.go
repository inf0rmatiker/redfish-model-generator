/* -----------------------------------------------------------------
* route_distinguisher_range.go -
*
* DMTF Redfish RouteDistinguisherRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Route Distinguisher (RD) number range for an Ethernet fabric.
type RouteDistinguisherRange struct {
	// Lower Route Distinguisher (RD) number.
	Lower int `json:"Lower,omitempty"`

	// Upper Route Distinguisher (RD) number.
	Upper int `json:"Upper,omitempty"`

}
