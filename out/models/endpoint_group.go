/* -----------------------------------------------------------------
* endpoint_group.go -
*
* DMTF Redfish EndpointGroup resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The EndpointGroup schema describes group of endpoints that are managed as a unit.
type EndpointGroup struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The access state for this group.
	AccessState AccessState `json:"AccessState,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The endpoints in this endpoint group.
	Endpoints array `json:"Endpoints,omitempty"`

	Endpoints@odata.count count `json:"Endpoints@odata.count,omitempty"`

	// The endpoint group type.
	GroupType GroupType `json:"GroupType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The durable name for the endpoint group.
	Identifier Identifier `json:"Identifier,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication if access to the resources through the endpoint group is preferred.
	Preferred bool `json:"Preferred,omitempty"`

	// The SCSI-defined identifier for this group.
	TargetEndpointGroupIdentifier int `json:"TargetEndpointGroupIdentifier,omitempty"`

}
