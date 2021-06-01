/* -----------------------------------------------------------------
* bgp_route.go -
*
* DMTF Redfish BGPRoute resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// BGPRoute - Border Gateway Protocol (BGP) route properties.
// Modeled after DMTF Redfish schema BGPRoute
type BGPRoute struct {
	// Advertise inactive route status.
	AdvertiseInactiveRoutesEnabled bool `json:"AdvertiseInactiveRoutesEnabled,omitempty"`

	// Route distance for external routes.
	DistanceExternal int `json:"DistanceExternal,omitempty"`

	// Route distance for internal routes.
	DistanceInternal int `json:"DistanceInternal,omitempty"`

	// Route distance for local routes.
	DistanceLocal int `json:"DistanceLocal,omitempty"`

	// Compare router id status.
	ExternalCompareRouterIdEnabled bool `json:"ExternalCompareRouterIdEnabled,omitempty"`

	// Route flap dampening status.
	FlapDampingEnabled bool `json:"FlapDampingEnabled,omitempty"`

	// Send default route status.
	SendDefaultRouteEnabled bool `json:"SendDefaultRouteEnabled,omitempty"`

}
