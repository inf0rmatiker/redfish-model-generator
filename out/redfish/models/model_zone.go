/* -----------------------------------------------------------------
* zone.go -
*
* DMTF Redfish Zone resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Zone - The Zone schema describes a simple fabric zone for a Redfish implementation.
// Modeled after DMTF Redfish schema Zone
type Zone struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// This property indicates whether routing within this zone is enabled.
	DefaultRoutingEnabled bool `json:"DefaultRoutingEnabled,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Indicates accessiblity of endpoints in this zone to endpoints outside of this zone.
	ExternalAccessibility ExternalAccessibility `json:"ExternalAccessibility,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The identifiers for this zone.
	Identifiers []Identifier `json:"Identifiers,omitempty"`

	// The links to Resources related to but not subordinate to this Resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The type of zone.
	ZoneType ZoneType `json:"ZoneType,omitempty"`

}
