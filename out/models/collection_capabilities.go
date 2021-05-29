/* -----------------------------------------------------------------
* collection_capabilities.go -
*
* DMTF Redfish CollectionCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes the capabilities of a collection.
type CollectionCapabilities struct {
	// The list of capabilities supported by this resource.
	Capabilities array `json:"Capabilities,omitempty"`

	// The maximum number of members allowed in this collection.
	MaxMembers int `json:"MaxMembers,omitempty"`

}
