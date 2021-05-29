/* -----------------------------------------------------------------
* gen_z.go -
*
* DMTF Redfish GenZ resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Gen-Z related properties for an address pool.
type GenZ struct {
	// The Access Key required for this address pool.
	AccessKey string `json:"AccessKey,omitempty"`

	// The maximum value for the Component Identifier (CID).
	MaxCID int `json:"MaxCID,omitempty"`

	// The maximum value for the Subnet Identifier (SID).
	MaxSID int `json:"MaxSID,omitempty"`

	// The minimum value for the Component Identifier (CID).
	MinCID int `json:"MinCID,omitempty"`

	// The minimum value for the Subnet Identifier (SID).
	MinSID int `json:"MinSID,omitempty"`

}
