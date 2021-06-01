/* -----------------------------------------------------------------
* data_storage_line_of_service.go -
*
* SNIA Swordfish DataStorageLineOfService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DataStorageLineOfService - Describe data storage and provisioning capabilities.
// Modeled after SNIA Swordfish schema DataStorageLineOfService
type DataStorageLineOfService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// Required access capabilities.
	AccessCapabilities []StorageAccessCapability `json:"AccessCapabilities,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// True implies compression or deduplication of storage.
	IsSpaceEfficient bool `json:"IsSpaceEfficient,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Provisioning policy for storage.
	ProvisioningPolicy map[string]interface{} `json:"ProvisioningPolicy,omitempty"`

	// Required minimum number of available capacity source resources.
	RecoverableCapacitySourceCount int `json:"RecoverableCapacitySourceCount,omitempty"`

	// Expectations for time to access the primary store after disaster recover.
	RecoveryTimeObjectives map[string]interface{} `json:"RecoveryTimeObjectives,omitempty"`

}
