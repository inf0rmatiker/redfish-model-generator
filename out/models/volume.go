/* -----------------------------------------------------------------
* volume.go -
*
* DMTF Redfish Volume resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Volume contains properties used to describe a volume, virtual disk, LUN, or other logical storage entity for any system.
type Volume struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// Supported IO access capabilities.
	AccessCapabilities array `json:"AccessCapabilities,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An array of references to StoragePools allocated from this Volume.
	AllocatedPools StoragePoolCollection `json:"AllocatedPools,omitempty"`

	// The size of the smallest addressable unit (Block) of this volume in bytes.
	BlockSizeBytes int `json:"BlockSizeBytes,omitempty"`

	// Capacity utilization.
	Capacity Capacity `json:"Capacity,omitempty"`

	// The size in bytes of this Volume.
	CapacityBytes int `json:"CapacityBytes,omitempty"`

	// An array of space allocations to this volume.
	CapacitySources array `json:"CapacitySources,omitempty"`

	CapacitySources@odata.count count `json:"CapacitySources@odata.count,omitempty"`

	// Indicator of whether or not the Volume has compression enabled.
	Compressed bool `json:"Compressed,omitempty"`

	// Indicator of whether or not the Volume has deduplication enabled.
	Deduplicated bool `json:"Deduplicated,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// A user-configurable string to name the volume.
	DisplayName string `json:"DisplayName,omitempty"`

	// Is this Volume encrypted.
	Encrypted bool `json:"Encrypted,omitempty"`

	// The types of encryption used by this Volume.
	EncryptionTypes array `json:"EncryptionTypes,omitempty"`

	// Indicates the IO performance mode setting for the volume.
	IOPerfModeEnabled bool `json:"IOPerfModeEnabled,omitempty"`

	// Statistics for this volume.
	IOStatistics IOStatistics `json:"IOStatistics,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The Durable names for the volume.
	Identifiers array `json:"Identifiers,omitempty"`

	// Indicates the Initialization Method used for this volume. If InitializeMethod is not specified, the InitializeMethod should be Foreground.
	InitializeMethod InitializeMethod `json:"InitializeMethod,omitempty"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// Indicates the host-visible LogicalUnitNumber assigned to this Volume.
	LogicalUnitNumber int `json:"LogicalUnitNumber,omitempty"`

	// Low space warning.
	LowSpaceWarningThresholdPercents []integer `json:"LowSpaceWarningThresholdPercents,omitempty"`

	// The manufacturer or OEM of this storage volume.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// Max Block size in bytes.
	MaxBlockSizeBytes int `json:"MaxBlockSizeBytes,omitempty"`

	// Indicates the number of media elements used per span in the secondary RAID for a hierarchical RAID type.
	MediaSpanCount int `json:"MediaSpanCount,omitempty"`

	// The model number for this storage volume.
	Model string `json:"Model,omitempty"`

	// This property contains properties to use when Volume is used to describe an NVMe Namespace.
	NVMeNamespaceProperties NVMeNamespaceProperties `json:"NVMeNamespaceProperties,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The operations currently running on the Volume.
	Operations array `json:"Operations,omitempty"`

	// The size in bytes of this Volume's optimum IO size.
	OptimumIOSizeBytes int `json:"OptimumIOSizeBytes,omitempty"`

	// This property specifies the volume's storage allocation, or provisioning policy.
	ProvisioningPolicy ProvisioningPolicy `json:"ProvisioningPolicy,omitempty"`

	// The RAID type of this volume.
	RAIDType RAIDType `json:"RAIDType,omitempty"`

	// Indicates the read cache policy setting for the Volume.
	ReadCachePolicy ReadCachePolicyType `json:"ReadCachePolicy,omitempty"`

	// Current number of capacity source resources that are available as replacements.
	RecoverableCapacitySourceCount int `json:"RecoverableCapacitySourceCount,omitempty"`

	// The percentage of the capacity remaining in the Volume.
	RemainingCapacityPercent int `json:"RemainingCapacityPercent,omitempty"`

	// Describes this storage volume in its role as a target replica.
	ReplicaInfo ReplicaInfo `json:"ReplicaInfo,omitempty"`

	// The resources that are target replicas of this source.
	ReplicaTargets []map[string]string `json:"ReplicaTargets,omitempty"`

	ReplicaTargets@odata.count count `json:"ReplicaTargets@odata.count,omitempty"`

	// The property contains the status of the Volume.
	Status Status `json:"Status,omitempty"`

	// An array of references to Storage Groups that includes this volume.
	StorageGroups StorageGroupCollection `json:"StorageGroups,omitempty"`

	// The number of blocks (bytes) in a strip in a disk array that uses striped data mapping.
	StripSizeBytes int `json:"StripSizeBytes,omitempty"`

	// The type of this volume.
	VolumeType VolumeType `json:"VolumeType,omitempty"`

	// Indicates the Volume usage type setting for the Volume.
	VolumeUsage VolumeUsageType `json:"VolumeUsage,omitempty"`

	// Indicates the write cache policy setting for the Volume.
	WriteCachePolicy WriteCachePolicyType `json:"WriteCachePolicy,omitempty"`

	// Indicates the WriteCacheState policy setting for the Volume.
	WriteCacheState WriteCacheStateType `json:"WriteCacheState,omitempty"`

	// The policy that the RAID volume is using to address the write hole issue.
	WriteHoleProtectionPolicy WriteHoleProtectionPolicyType `json:"WriteHoleProtectionPolicy,omitempty"`

}
