/* -----------------------------------------------------------------
* consistency_group.go -
*
* SNIA Swordfish ConsistencyGroup resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ConsistencyGroup - A collection of volumes grouped together to ensure write order consistency across all those volumes.
// Modeled after SNIA Swordfish schema ConsistencyGroup
type ConsistencyGroup struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The consistency method used by this group.
	ConsistencyMethod map[string]interface{} `json:"ConsistencyMethod,omitempty"`

	// The consistency type used by this group.
	ConsistencyType map[string]interface{} `json:"ConsistencyType,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// This value is true when the consistency group is in a consistent state.
	IsConsistent bool `json:"IsConsistent,omitempty"`

	// Contains links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Describes this storage group in its role as a target for replication.
	ReplicaInfo map[string]interface{} `json:"ReplicaInfo,omitempty"`

	// The resources that are target replicas of this source.
	ReplicaTargets []idRef `json:"ReplicaTargets,omitempty"`

	ReplicaTargetsOdataCount int `json:"ReplicaTargets@odata.count,omitempty"`

	// The property contains the status of the ConsistencyGroup.
	Status map[string]interface{} `json:"Status,omitempty"`

	// Volumes in this storage group.
	Volumes []Volume `json:"Volumes,omitempty"`

	VolumesOdataCount int `json:"Volumes@odata.count,omitempty"`

}
