/* -----------------------------------------------------------------
* volume.go -
*
* SNIA Swordfish Volume resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Volume - Volume contains properties used to describe a volume, virtual disk, LUN, or other logical storage entity for any system.
// Modeled after SNIA Swordfish schema Volume
type Volume struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// Supported IO access capabilities.
	AccessCapabilities []StorageAccessCapability `json:"AccessCapabilities,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An array of references to StoragePools allocated from this Volume.
	AllocatedPools map[string]interface{} `json:"AllocatedPools,omitempty"`

	// The size of the smallest addressable unit (Block) of this volume in bytes.
	BlockSizeBytes int `json:"BlockSizeBytes,omitempty"`

	// Capacity utilization.
	Capacity map[string]interface{} `json:"Capacity,omitempty"`

	// The size in bytes of this Volume.
	CapacityBytes int `json:"CapacityBytes,omitempty"`

	// An array of space allocations to this volume.
	CapacitySources []CapacitySource `json:"CapacitySources,omitempty"`

	CapacitySourcesOdataCount int `json:"CapacitySources@odata.count,omitempty"`

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
	EncryptionTypes []EncryptionTypes `json:"EncryptionTypes,omitempty"`

	// Indicates the IO performance mode setting for the volume.
	IOPerfModeEnabled bool `json:"IOPerfModeEnabled,omitempty"`

	// Statistics for this volume.
	IOStatistics map[string]interface{} `json:"IOStatistics,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The Durable names for the volume.
	Identifiers []Identifier `json:"Identifiers,omitempty"`

	// Indicates the Initialization Method used for this volume. If InitializeMethod is not specified, the InitializeMethod should be Foreground.
	InitializeMethod map[string]interface{} `json:"InitializeMethod,omitempty"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// Indicates the host-visible LogicalUnitNumber assigned to this Volume.
	LogicalUnitNumber int `json:"LogicalUnitNumber,omitempty"`

	// Low space warning.
	LowSpaceWarningThresholdPercents []int `json:"LowSpaceWarningThresholdPercents,omitempty"`

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
	Operations []Operation `json:"Operations,omitempty"`

	// The size in bytes of this Volume's optimum IO size.
	OptimumIOSizeBytes int `json:"OptimumIOSizeBytes,omitempty"`

	// This property specifies the volume's storage allocation, or provisioning policy.
	ProvisioningPolicy map[string]interface{} `json:"ProvisioningPolicy,omitempty"`

	// The RAID type of this volume.
	RAIDType map[string]interface{} `json:"RAIDType,omitempty"`

	// Indicates the read cache policy setting for the Volume.
	ReadCachePolicy map[string]interface{} `json:"ReadCachePolicy,omitempty"`

	// Current number of capacity source resources that are available as replacements.
	RecoverableCapacitySourceCount int `json:"RecoverableCapacitySourceCount,omitempty"`

	// The percentage of the capacity remaining in the Volume.
	RemainingCapacityPercent int `json:"RemainingCapacityPercent,omitempty"`

	// Describes this storage volume in its role as a target replica.
	ReplicaInfo map[string]interface{} `json:"ReplicaInfo,omitempty"`

	// The resources that are target replicas of this source.
	ReplicaTargets []idRef `json:"ReplicaTargets,omitempty"`

	ReplicaTargetsOdataCount int `json:"ReplicaTargets@odata.count,omitempty"`

	// The property contains the status of the Volume.
	Status map[string]interface{} `json:"Status,omitempty"`

	// An array of references to Storage Groups that includes this volume.
	StorageGroups map[string]interface{} `json:"StorageGroups,omitempty"`

	// The number of blocks (bytes) in a strip in a disk array that uses striped data mapping.
	StripSizeBytes int `json:"StripSizeBytes,omitempty"`

	// The type of this volume.
	VolumeType map[string]interface{} `json:"VolumeType,omitempty"`

	// Indicates the Volume usage type setting for the Volume.
	VolumeUsage map[string]interface{} `json:"VolumeUsage,omitempty"`

	// Indicates the write cache policy setting for the Volume.
	WriteCachePolicy map[string]interface{} `json:"WriteCachePolicy,omitempty"`

	// Indicates the WriteCacheState policy setting for the Volume.
	WriteCacheState map[string]interface{} `json:"WriteCacheState,omitempty"`

	// The policy that the RAID volume is using to address the write hole issue.
	WriteHoleProtectionPolicy map[string]interface{} `json:"WriteHoleProtectionPolicy,omitempty"`

}
