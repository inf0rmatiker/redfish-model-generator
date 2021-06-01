/* -----------------------------------------------------------------
* vlan.go -
*
* DMTF Redfish VLAN resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// This indicates if this VLAN is enabled.
	VLANEnable bool `json:"VLANEnable,omitempty"`

	// This indicates the VLAN identifier for this VLAN.
	VLANId VLANId `json:"VLANId,omitempty"`

}
