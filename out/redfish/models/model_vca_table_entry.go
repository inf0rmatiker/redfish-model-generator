/* -----------------------------------------------------------------
* vca_table_entry.go -
*
* DMTF Redfish VCATableEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VCATableEntry - The Virtual Channel Action Table entry corresponding to a specific Virtual Channel.
// Modeled after DMTF Redfish schema VCATableEntry
type VCATableEntry struct {
	// The configured threshold.
	Threshold string `json:"Threshold,omitempty"`

	// The bits corresponding to the supported Virtual Channel.
	VCMask string `json:"VCMask,omitempty"`

}
