/* -----------------------------------------------------------------
* class_of_service_collection.go -
*
* SNIA Swordfish ClassOfServiceCollection resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ClassOfServiceCollection - A Collection of ClassOfService or LineofService resource instances.
// Modeled after SNIA Swordfish schema ClassOfServiceCollection
type ClassOfServiceCollection struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of this collection.
	Members []LineOfService `json:"Members"`

	MembersOdataCount int `json:"Members@odata.count"`

	MembersOdataNextLink map[string]interface{} `json:"Members@odata.nextLink,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
