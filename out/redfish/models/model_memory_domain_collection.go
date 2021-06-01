/* -----------------------------------------------------------------
* memory_domain_collection.go -
*
* DMTF Redfish MemoryDomainCollection resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MemoryDomainCollection - The collection of MemoryDomain Resource instances.
// Modeled after DMTF Redfish schema MemoryDomainCollection
type MemoryDomainCollection struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of this collection.
	Members []MemoryDomain `json:"Members"`

	MembersOdataCount int `json:"Members@odata.count"`

	MembersOdataNextLink map[string]interface{} `json:"Members@odata.nextLink,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
