/* -----------------------------------------------------------------
* data_storage_lo_s_capabilities.go -
*
* SNIA Swordfish DataStorageLoSCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DataStorageLoSCapabilities - Describe data storage and provisioning capabilities.
// Modeled after SNIA Swordfish schema DataStorageLoSCapabilities
type DataStorageLoSCapabilities struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// Maximum number of capacity source resources for the purpose of recovery from a failure.
	MaximumRecoverableCapacitySourceCount int `json:"MaximumRecoverableCapacitySourceCount,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Supported access capabilities.
	SupportedAccessCapabilities []StorageAccessCapability `json:"SupportedAccessCapabilities,omitempty"`

	// Collection of known and supported DataStorageLinesOfService.
	SupportedLinesOfService []DataStorageLineOfService `json:"SupportedLinesOfService,omitempty"`

	SupportedLinesOfServiceOdataCount int `json:"SupportedLinesOfService@odata.count,omitempty"`

	// Thin allows over allocation of storage.
	SupportedProvisioningPolicies []ProvisioningPolicy `json:"SupportedProvisioningPolicies,omitempty"`

	// Supported expectations for time to access the primary store after recovery.
	SupportedRecoveryTimeObjectives []RecoveryAccessScope `json:"SupportedRecoveryTimeObjectives,omitempty"`

	// Allows compression or deduplication of storage.
	SupportsSpaceEfficiency bool `json:"SupportsSpaceEfficiency,omitempty"`

}
