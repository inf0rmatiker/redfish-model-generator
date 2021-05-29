/* -----------------------------------------------------------------
* vlan.go -
*
* DMTF Redfish VLAN resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The attributes of a VLAN.
type VLAN struct {
	// An indication of whether this VLAN is enabled for this VLAN network interface.
	VLANEnable bool `json:"VLANEnable,omitempty"`

	// The ID for this VLAN.
	VLANId VLANId `json:"VLANId,omitempty"`

	// The priority for this VLAN.
	VLANPriority VLANPriority `json:"VLANPriority,omitempty"`

}
