/* -----------------------------------------------------------------
* service_root.go -
*
* DMTF Redfish ServiceRoot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ServiceRoot schema describes the root of the Redfish Service, located at the '/redfish/v1' URI.  All other Resources accessible through the Redfish interface on this device are linked directly or indirectly from the Service Root.
type ServiceRoot struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The link to the Account Service.
	AccountService AccountService `json:"AccountService,omitempty"`

	// The link to the aggregation service.
	AggregationService AggregationService `json:"AggregationService,omitempty"`

	// The link to the Certificate Service.
	CertificateService CertificateService `json:"CertificateService,omitempty"`

	// The link to a collection of chassis.
	Chassis ChassisCollection `json:"Chassis,omitempty"`

	// The link to the Composition Service.
	CompositionService CompositionService `json:"CompositionService,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to the Event Service.
	EventService EventService `json:"EventService,omitempty"`

	// The link to a collection of all fabric entities.
	Fabrics FabricCollection `json:"Fabrics,omitempty"`

	// The link to a collection of facilities.
	Facilities FacilityCollection `json:"Facilities,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The link to the JobService.
	JobService JobService `json:"JobService,omitempty"`

	// The link to a collection of JSON Schema files.
	JsonSchemas JsonSchemaFileCollection `json:"JsonSchemas,omitempty"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links"`

	// The link to a collection of managers.
	Managers ManagerCollection `json:"Managers,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to a set of power equipment.
	PowerEquipment PowerEquipment `json:"PowerEquipment,omitempty"`

	// The product associated with this Redfish Service.
	Product string `json:"Product,omitempty"`

	// The information about protocol features that the service supports.
	ProtocolFeaturesSupported ProtocolFeaturesSupported `json:"ProtocolFeaturesSupported,omitempty"`

	// The version of the Redfish Service.
	RedfishVersion string `json:"RedfishVersion,omitempty"`

	// The link to a collection of Registries.
	Registries MessageRegistryFileCollection `json:"Registries,omitempty"`

	// The link to a collection of all Resource Block Resources.  This collection is intended for implementations that do not contain a Composition Service but that expose Resources to an orchestrator that implements a Composition Service.
	ResourceBlocks ResourceBlockCollection `json:"ResourceBlocks,omitempty"`

	// The link to the Sessions Service.
	SessionService SessionService `json:"SessionService,omitempty"`

	// The link to a collection of storage subsystems.
	Storage StorageCollection `json:"Storage,omitempty"`

	// The link to a collection of all storage service entities.
	StorageServices StorageServiceCollection `json:"StorageServices,omitempty"`

	// The link to a collection of storage systems.
	StorageSystems StorageSystemCollection `json:"StorageSystems,omitempty"`

	// The link to a collection of systems.
	Systems ComputerSystemCollection `json:"Systems,omitempty"`

	// The link to the Task Service.
	Tasks TaskService `json:"Tasks,omitempty"`

	// The link to the Telemetry Service.
	TelemetryService TelemetryService `json:"TelemetryService,omitempty"`

	// Unique identifier for a service instance.  When SSDP is used, this value should be an exact match of the UUID value returned in a 200 OK from an SSDP M-SEARCH request during discovery.
	UUID UUID `json:"UUID,omitempty"`

	// The link to the Update Service.
	UpdateService UpdateService `json:"UpdateService,omitempty"`

	// The vendor or manufacturer associated with this Redfish Service.
	Vendor string `json:"Vendor,omitempty"`

}
