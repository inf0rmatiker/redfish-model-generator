/* -----------------------------------------------------------------
* vcat_entry_collection.go -
*
* DMTF Redfish VCATEntryCollection resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VCATEntryCollection - The collection of VCATEntry Resource instances.
// Modeled after DMTF Redfish schema VCATEntryCollection
type VCATEntryCollection struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of this collection.
	Members []VCATEntry `json:"Members"`

	MembersOdataCount int `json:"Members@odata.count"`

	MembersOdataNextLink map[string]interface{} `json:"Members@odata.nextLink,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
