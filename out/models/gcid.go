/* -----------------------------------------------------------------
* gcid.go -
*
* DMTF Redfish GCID resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Global Component ID (GCID).
type GCID struct {
	// The component identifier portion of the GCID for the entity.
	CID string `json:"CID,omitempty"`

	// The subnet identifier portion of the GCID for the entity.
	SID string `json:"SID,omitempty"`

}
