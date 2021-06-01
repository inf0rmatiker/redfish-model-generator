/* -----------------------------------------------------------------
* service_root.go -
*
* DMTF Redfish ServiceRoot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ServiceRoot - This object represents the root Redfish service.
// Modeled after DMTF Redfish schema ServiceRoot
type ServiceRoot struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// This is a link to the Account Service.
	AccountService map[string]interface{} `json:"AccountService,omitempty"`

	// This is a link to a collection of Chassis.
	Chassis map[string]interface{} `json:"Chassis,omitempty"`

	// This is a link to the CompositionService.
	CompositionService map[string]interface{} `json:"CompositionService,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// This is a link to the EventService.
	EventService map[string]interface{} `json:"EventService,omitempty"`

	// A link to a collection of all fabric entities.
	Fabrics map[string]interface{} `json:"Fabrics,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// This is a link to a collection of Json-Schema files.
	JsonSchemas map[string]interface{} `json:"JsonSchemas,omitempty"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links"`

	// This is a link to a collection of Managers.
	Managers map[string]interface{} `json:"Managers,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The product associated with this Redfish service.
	Product string `json:"Product,omitempty"`

	// Contains information about protocol features supported by the service.
	ProtocolFeaturesSupported ProtocolFeaturesSupported `json:"ProtocolFeaturesSupported,omitempty"`

	// The version of the Redfish service.
	RedfishVersion string `json:"RedfishVersion,omitempty"`

	// This is a link to a collection of Registries.
	Registries map[string]interface{} `json:"Registries,omitempty"`

	// This is a link to the Sessions Service.
	SessionService map[string]interface{} `json:"SessionService,omitempty"`

	// A link to a collection of all storage service entities.
	StorageServices map[string]interface{} `json:"StorageServices,omitempty"`

	// This is a link to a collection of storage systems.
	StorageSystems map[string]interface{} `json:"StorageSystems,omitempty"`

	// This is a link to a collection of Systems.
	Systems map[string]interface{} `json:"Systems,omitempty"`

	// This is a link to the Task Service.
	Tasks map[string]interface{} `json:"Tasks,omitempty"`

	// Unique identifier for a service instance. When SSDP is used, this value should be an exact match of the UUID value returned in a 200OK from an SSDP M-SEARCH request during discovery.
	UUID map[string]interface{} `json:"UUID,omitempty"`

	// This is a link to the UpdateService.
	UpdateService map[string]interface{} `json:"UpdateService,omitempty"`

}
