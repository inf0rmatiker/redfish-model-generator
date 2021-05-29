/* -----------------------------------------------------------------
* infini_band.go -
*
* DMTF Redfish InfiniBand resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes InfiniBand capabilities, status, and configuration of a network device function.
type InfiniBand struct {
	// The maximum transmission unit (MTU) configured for this network device function.
	MTUSize int `json:"MTUSize,omitempty"`

	// This is the currently configured node GUID of the network device function.
	NodeGUID string `json:"NodeGUID,omitempty"`

	// The permanent node GUID assigned to this network device function.
	PermanentNodeGUID string `json:"PermanentNodeGUID,omitempty"`

	// The permanent port GUID assigned to this network device function.
	PermanentPortGUID string `json:"PermanentPortGUID,omitempty"`

	// The permanent system GUID assigned to this network device function.
	PermanentSystemGUID string `json:"PermanentSystemGUID,omitempty"`

	// The currently configured port GUID of the network device function.
	PortGUID string `json:"PortGUID,omitempty"`

	// The maximum transmission unit (MTU) sizes supported for this network device function.
	SupportedMTUSizes []integer `json:"SupportedMTUSizes,omitempty"`

	// This is the currently configured system GUID of the network device function.
	SystemGUID string `json:"SystemGUID,omitempty"`

}
