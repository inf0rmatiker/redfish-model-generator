/* -----------------------------------------------------------------
* endpoint_group.go -
*
* SNIA Swordfish EndpointGroup resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EndpointGroup - A group of endpoints that are managed as a unit.
// Modeled after SNIA Swordfish schema EndpointGroup
type EndpointGroup struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// Access State for this group.
	AccessState map[string]interface{} `json:"AccessState,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The Endpoints.
	Endpoints []Endpoint `json:"Endpoints,omitempty"`

	EndpointsOdataCount int `json:"Endpoints@odata.count,omitempty"`

	// Endpoint group type.
	GroupType GroupType `json:"GroupType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// Contains links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Access to resource is preferred.
	Preferred bool `json:"Preferred,omitempty"`

	// A defined identifier for this group.
	TargetEndpointGroupIdentifier int `json:"TargetEndpointGroupIdentifier,omitempty"`

}
