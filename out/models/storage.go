/* -----------------------------------------------------------------
* storage.go -
*
* DMTF Redfish Storage resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Storage schema defines a storage subsystem and its respective properties.  A storage subsystem represents a set of physical or virtual storage controllers and the resources, such as volumes, that can be accessed from that subsystem.
type Storage struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The consistency groups, each of which contains a set of volumes that are treated by an application or set of applications as a single resource, that are managed by this storage subsystem.
	ConsistencyGroups ConsistencyGroupCollection `json:"ConsistencyGroups,omitempty"`

	// The set of controllers instantiated by this storage subsystem.
	Controllers StorageControllerCollection `json:"Controllers,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The set of drives attached to the storage controllers that this resource represents.
	Drives array `json:"Drives,omitempty"`

	Drives@odata.count count `json:"Drives@odata.count,omitempty"`

	// All of the endpoint groups, each of which contains a set of endpoints that are used for a common purpose such as an ACL or logical identification, that belong to this storage subsystem.
	EndpointGroups EndpointGroupCollection `json:"EndpointGroups,omitempty"`

	// All file systems that are allocated by this storage subsystem.
	FileSystems FileSystemCollection `json:"FileSystems,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The durable names for the storage subsystem.
	Identifiers array `json:"Identifiers,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Redundancy information for the storage subsystem.
	Redundancy array `json:"Redundancy,omitempty"`

	Redundancy@odata.count count `json:"Redundancy@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The set of storage controllers that this resource represents.
	StorageControllers array `json:"StorageControllers,omitempty"`

	StorageControllers@odata.count count `json:"StorageControllers@odata.count,omitempty"`

	// All of the storage groups, each of which contains a set of volumes and endpoints that are managed as a group for mapping and masking, that belong to this storage subsystem.
	StorageGroups StorageGroupCollection `json:"StorageGroups,omitempty"`

	// The set of all storage pools that are allocated by this storage subsystem.  A storage pool is the set of storage capacity that can be used to produce volumes or other storage pools.
	StoragePools StoragePoolCollection `json:"StoragePools,omitempty"`

	// The set of volumes that the storage controllers produce.
	Volumes VolumeCollection `json:"Volumes,omitempty"`

}
