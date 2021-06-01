/* -----------------------------------------------------------------
* resource_block_collection.go -
*
* DMTF Redfish ResourceBlockCollection resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ResourceBlockCollection - The collection of ResourceBlock resource instances.
// Modeled after DMTF Redfish schema ResourceBlockCollection
type ResourceBlockCollection struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of this collection.
	Members []ResourceBlock `json:"Members"`

	MembersOdataCount int `json:"Members@odata.count"`

	MembersOdataNextLink map[string]interface{} `json:"Members@odata.nextLink,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
