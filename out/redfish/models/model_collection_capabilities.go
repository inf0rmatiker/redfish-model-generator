/* -----------------------------------------------------------------
* collection_capabilities.go -
*
* DMTF Redfish CollectionCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CollectionCapabilities - This type describes the capabilities of a collection.
// Modeled after DMTF Redfish schema CollectionCapabilities
type CollectionCapabilities struct {
	// This property contains the list of capabilities supported by this resource.
	Capabilities []Capability `json:"Capabilities,omitempty"`

}
