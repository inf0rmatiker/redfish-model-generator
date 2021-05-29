/* -----------------------------------------------------------------
* connection_method.go -
*
* DMTF Redfish ConnectionMethod resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ConnectionMethod schema describes the protocol, provider, or other method used to communicate to a given access point for a Redfish aggregation service.
type ConnectionMethod struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The type of connection method.
	ConnectionMethodType ConnectionMethodType `json:"ConnectionMethodType,omitempty"`

	// The variant of connection method.
	ConnectionMethodVariant string `json:"ConnectionMethodVariant,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
