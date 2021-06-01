/* -----------------------------------------------------------------
* file_system.go -
*
* SNIA Swordfish FileSystem resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// FileSystem - An instance of a hierarchical namespace of files.
// Modeled after SNIA Swordfish schema FileSystem
type FileSystem struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// An array of supported IO access capabilities.
	AccessCapabilities []StorageAccessCapability `json:"AccessCapabilities,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Block size of the file system in bytes.
	BlockSizeBytes int `json:"BlockSizeBytes,omitempty"`

	// Capacity allocated to the file system.
	Capacity map[string]interface{} `json:"Capacity,omitempty"`

	// An array of capacity sources for the file system.
	CapacitySources []CapacitySource `json:"CapacitySources,omitempty"`

	CapacitySourcesOdataCount int `json:"CapacitySources@odata.count,omitempty"`

	// The case of file names is preserved by the file system.
	CasePreserved bool `json:"CasePreserved,omitempty"`

	// Case sensitive file names are supported by the file system.
	CaseSensitive bool `json:"CaseSensitive,omitempty"`

	// An array of the character sets or encodings supported by the file system.
	CharacterCodeSet []CharacterCodeSet `json:"CharacterCodeSet,omitempty"`

	// A value indicating the minimum file allocation size imposed by the file system.
	ClusterSizeBytes int `json:"ClusterSizeBytes,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An array of exported file shares of this file system.
	ExportedShares map[string]interface{} `json:"ExportedShares,omitempty"`

	// Statistics for this FileSystem.
	IOStatistics map[string]interface{} `json:"IOStatistics,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The durable names for this file system.
	Identifiers []Identifier `json:"Identifiers,omitempty"`

	// An array of imported file shares.
	ImportedShares []ImportedShare `json:"ImportedShares,omitempty"`

	// Contains links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An array of low space warning threshold percentages for the file system.
	LowSpaceWarningThresholdPercents []int `json:"LowSpaceWarningThresholdPercents,omitempty"`

	// A value indicating the maximum length of a file name within the file system.
	MaxFileNameLengthBytes int `json:"MaxFileNameLengthBytes,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Current number of capacity source resources that are available as replacements.
	RecoverableCapacitySourceCount int `json:"RecoverableCapacitySourceCount,omitempty"`

	// Remaining capacity allocated to the file system.
	RemainingCapacity map[string]interface{} `json:"RemainingCapacity,omitempty"`

	// The percentage of the capacity remaining in the FileSystem.
	RemainingCapacityPercent int `json:"RemainingCapacityPercent,omitempty"`

	// This value describes the replica attributes if this file system is a replica.
	ReplicaInfo map[string]interface{} `json:"ReplicaInfo,omitempty"`

	// The resources that are target replicas of this source.
	ReplicaTargets []idRef `json:"ReplicaTargets,omitempty"`

	ReplicaTargetsOdataCount int `json:"ReplicaTargets@odata.count,omitempty"`

}
