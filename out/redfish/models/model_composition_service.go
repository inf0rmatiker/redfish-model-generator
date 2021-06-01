/* -----------------------------------------------------------------
* composition_service.go -
*
* DMTF Redfish CompositionService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CompositionService - The CompositionService schema describes a composition service and its properties and links to the resources available for composition.
// Modeled after DMTF Redfish schema CompositionService
type CompositionService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An indication of whether this service is allowed to overprovision a composition relative to the composition request.
	AllowOverprovisioning bool `json:"AllowOverprovisioning,omitempty"`

	// An indication of whether a client can request that a specific resource zone fulfill a composition request.
	AllowZoneAffinity bool `json:"AllowZoneAffinity,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The resource blocks available on the service.
	ResourceBlocks map[string]interface{} `json:"ResourceBlocks,omitempty"`

	// The resource zones available on the service.
	ResourceZones map[string]interface{} `json:"ResourceZones,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
