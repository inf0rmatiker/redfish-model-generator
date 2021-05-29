/* -----------------------------------------------------------------
* zone.go -
*
* DMTF Redfish Zone resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Zone schema describes a simple fabric zone for a Redfish implementation.
type Zone struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// This property indicates whether routing within this zone is enabled.
	DefaultRoutingEnabled bool `json:"DefaultRoutingEnabled,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Indicates accessibility of endpoints in this zone to endpoints outside of this zone.
	ExternalAccessibility ExternalAccessibility `json:"ExternalAccessibility,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The durable names for the zone.
	Identifiers array `json:"Identifiers,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The type of zone.
	ZoneType ZoneType `json:"ZoneType,omitempty"`

}
