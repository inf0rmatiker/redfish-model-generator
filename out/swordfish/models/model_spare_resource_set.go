/* -----------------------------------------------------------------
* spare_resource_set.go -
*
* SNIA Swordfish SpareResourceSet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SpareResourceSet - A description of a set of spare resources.
// Modeled after SNIA Swordfish schema SpareResourceSet
type SpareResourceSet struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Contains links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Location where this set of spares is kept.
	OnHandLocation map[string]interface{} `json:"OnHandLocation,omitempty"`

	// This set is available online.
	OnLine bool `json:"OnLine,omitempty"`

	// The type of resources in the set.
	ResourceType string `json:"ResourceType,omitempty"`

	// Amount of time needed to make an on-hand resource available as a spare.
	TimeToProvision string `json:"TimeToProvision,omitempty"`

	// Amount of time needed to get more on-hand resources.
	TimeToReplenish string `json:"TimeToReplenish,omitempty"`

}
