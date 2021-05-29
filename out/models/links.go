/* -----------------------------------------------------------------
* links.go -
*
* DMTF Redfish Links resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Add ability to manage spare capacity.
type Links struct {
	// A pointer to the data volumes this volume serves as a cache volume.
	CacheDataVolumes array `json:"CacheDataVolumes,omitempty"`

	CacheDataVolumes@odata.count count `json:"CacheDataVolumes@odata.count,omitempty"`

	// A pointer to the cache volume source for this volume.
	CacheVolumeSource Volume `json:"CacheVolumeSource,omitempty"`

	// The ClassOfService that this storage volume conforms to.
	ClassOfService ClassOfService `json:"ClassOfService,omitempty"`

	// An array of references to the client Endpoints associated with this volume.
	ClientEndpoints array `json:"ClientEndpoints,omitempty"`

	ClientEndpoints@odata.count count `json:"ClientEndpoints@odata.count,omitempty"`

	// An array of references to the ConsistencyGroups associated with this volume.
	ConsistencyGroups array `json:"ConsistencyGroups,omitempty"`

	ConsistencyGroups@odata.count count `json:"ConsistencyGroups@odata.count,omitempty"`

	// An array of references to the drives which are dedicated spares for this volume.
	DedicatedSpareDrives array `json:"DedicatedSpareDrives,omitempty"`

	DedicatedSpareDrives@odata.count count `json:"DedicatedSpareDrives@odata.count,omitempty"`

	// An array of references to the drives which contain this volume. This will reference Drives that either wholly or only partly contain this volume.
	Drives array `json:"Drives,omitempty"`

	Drives@odata.count count `json:"Drives@odata.count,omitempty"`

	// A pointer to the Resource that serves as a journaling media for this volume.
	JournalingMedia Resource `json:"JournalingMedia,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A pointer to the Storage resource that owns or contains this volume.
	OwningStorageResource Storage `json:"OwningStorageResource,omitempty"`

	// A pointer to the StorageService that owns or contains this volume.
	OwningStorageService StorageService `json:"OwningStorageService,omitempty"`

	// An array of references to the server Endpoints associated with this volume.
	ServerEndpoints array `json:"ServerEndpoints,omitempty"`

	ServerEndpoints@odata.count count `json:"ServerEndpoints@odata.count,omitempty"`

	// An array of references to SpareResourceSets.
	SpareResourceSets array `json:"SpareResourceSets,omitempty"`

	SpareResourceSets@odata.count count `json:"SpareResourceSets@odata.count,omitempty"`

	// An array of references to the StorageGroups associated with this volume.
	StorageGroups array `json:"StorageGroups,omitempty"`

	StorageGroups@odata.count count `json:"StorageGroups@odata.count,omitempty"`

}
