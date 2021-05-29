/* -----------------------------------------------------------------
* vca_table_entry.go -
*
* DMTF Redfish VCATableEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Virtual Channel Action Table entry corresponding to a specific Virtual Channel.
type VCATableEntry struct {
	// The configured threshold.
	Threshold string `json:"Threshold,omitempty"`

	// The bits corresponding to the supported Virtual Channel.
	VCMask string `json:"VCMask,omitempty"`

}
