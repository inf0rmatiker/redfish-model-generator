/* -----------------------------------------------------------------
* storage_group.go -
*
* SNIA Swordfish StorageGroup resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// StorageGroup - Collection of resources that are managed and exposed to hosts as a group.
// Modeled after SNIA Swordfish schema StorageGroup
type StorageGroup struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// AccessState for this storage group.
	AccessState map[string]interface{} `json:"AccessState,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The Authentication method used for the Endpoints involved in this StorageGroup.
	AuthenticationMethod AuthenticationMethod `json:"AuthenticationMethod,omitempty"`

	// The credential information used to authenticate the endpoints in this StorageGroup.
	ChapInfo []CHAPInformation `json:"ChapInfo,omitempty"`

	// Groups of client endpoints in this storage group.
	ClientEndpointGroups []EndpointGroup `json:"ClientEndpointGroups,omitempty"`

	ClientEndpointGroupsOdataCount int `json:"ClientEndpointGroups@odata.count,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// Contains links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// Mapped Volumes in this storage group.
	MappedVolumes []MappedVolume `json:"MappedVolumes,omitempty"`

	// Members are kept in a consistent state.
	MembersAreConsistent bool `json:"MembersAreConsistent,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Describes this storage group in its role as a target for replication.
	ReplicaInfo map[string]interface{} `json:"ReplicaInfo,omitempty"`

	// The resources that are target replicas of this source.
	ReplicaTargets []idRef `json:"ReplicaTargets,omitempty"`

	ReplicaTargetsOdataCount int `json:"ReplicaTargets@odata.count,omitempty"`

	// Groups of server endpoints in this storage group.
	ServerEndpointGroups []EndpointGroup `json:"ServerEndpointGroups,omitempty"`

	ServerEndpointGroupsOdataCount int `json:"ServerEndpointGroups@odata.count,omitempty"`

	// The property contains the status of the StorageGroup.
	Status map[string]interface{} `json:"Status,omitempty"`

	// Volumes in this storage group.
	Volumes []Volume `json:"Volumes,omitempty"`

	VolumesOdataCount int `json:"Volumes@odata.count,omitempty"`

	// Storage volumes are exposed to paths defined by the client and server endpoints.
	VolumesAreExposed bool `json:"VolumesAreExposed,omitempty"`

}
