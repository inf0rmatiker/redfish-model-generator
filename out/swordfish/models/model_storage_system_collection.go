/* -----------------------------------------------------------------
* storage_system_collection.go -
*
* SNIA Swordfish StorageSystemCollection resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// StorageSystemCollection - A Collection of ComputerSystem resource instances that each have an entry in HostingRoles set to 'StorageServer'.
// Modeled after SNIA Swordfish schema StorageSystemCollection
type StorageSystemCollection struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of this collection.
	Members []ComputerSystem `json:"Members"`

	MembersOdataCount int `json:"Members@odata.count"`

	MembersOdataNextLink map[string]interface{} `json:"Members@odata.nextLink,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
