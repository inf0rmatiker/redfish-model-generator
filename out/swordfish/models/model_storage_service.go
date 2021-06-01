/* -----------------------------------------------------------------
* storage_service.go -
*
* SNIA Swordfish StorageService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// StorageService - Collection of resources that are managed and exposed to hosts as a group.
// Modeled after SNIA Swordfish schema StorageService
type StorageService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The ClassesOfService that all storage in this StorageService can support.
	ClassesOfService map[string]interface{} `json:"ClassesOfService,omitempty"`

	// Client endpoint groups.
	ClientEndpointGroups map[string]interface{} `json:"ClientEndpointGroups,omitempty"`

	// The data protection capabilities of this service.
	DataProtectionLoSCapabilities map[string]interface{} `json:"DataProtectionLoSCapabilities,omitempty"`

	// The data security capabilities of this service.
	DataSecurityLoSCapabilities map[string]interface{} `json:"DataSecurityLoSCapabilities,omitempty"`

	// The data storage capabilities of this service.
	DataStorageLoSCapabilities map[string]interface{} `json:"DataStorageLoSCapabilities,omitempty"`

	// The default class of service for entities allocated by this storage service.
	DefaultClassOfService map[string]interface{} `json:"DefaultClassOfService,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The set of drives managed by this storage service.
	Drives map[string]interface{} `json:"Drives,omitempty"`

	// Client and Server endpoint groups.
	EndpointGroups map[string]interface{} `json:"EndpointGroups,omitempty"`

	// Endpoints.
	Endpoints map[string]interface{} `json:"Endpoints,omitempty"`

	// FileSystems.
	FileSystems map[string]interface{} `json:"FileSystems,omitempty"`

	// The IO connectivity capabilities of this service.
	IOConnectivityLoSCapabilities map[string]interface{} `json:"IOConnectivityLoSCapabilities,omitempty"`

	// The IO performance capabilities of this service.
	IOPerformanceLoSCapabilities map[string]interface{} `json:"IOPerformanceLoSCapabilities,omitempty"`

	// Statistics for this StorageService.
	IOStatistics map[string]interface{} `json:"IOStatistics,omitempty"`

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

	// Redundancy information for the storage subsystem.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	// Server endpoint groups.
	ServerEndpointGroups map[string]interface{} `json:"ServerEndpointGroups,omitempty"`

	// An array of SpareResourceSets.
	SpareResourceSets []SpareResourceSet `json:"SpareResourceSets,omitempty"`

	SpareResourceSetsOdataCount int `json:"SpareResourceSets@odata.count,omitempty"`

	// The property contains the status of the StorageService.
	Status map[string]interface{} `json:"Status,omitempty"`

	// StorageGroups.
	StorageGroups map[string]interface{} `json:"StorageGroups,omitempty"`

	// StoragePools.
	StoragePools map[string]interface{} `json:"StoragePools,omitempty"`

	// A reference to storage subsystems managed by this storage service.
	StorageSubsystems map[string]interface{} `json:"StorageSubsystems,omitempty"`

	// Volumes.
	Volumes map[string]interface{} `json:"Volumes,omitempty"`

}
