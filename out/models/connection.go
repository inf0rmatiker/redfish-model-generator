/* -----------------------------------------------------------------
* connection.go -
*
* DMTF Redfish Connection resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Connection schema describes the access permissions endpoints, or groups of endpoints, have with other resources in the service.
type Connection struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The permission keys required to access the specified resources for this connection.
	ConnectionKeys ConnectionKey `json:"ConnectionKeys,omitempty"`

	// The type of resources this connection specifies.
	ConnectionType ConnectionType `json:"ConnectionType,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The set of memory chunks and access capabilities specified for this connection.
	MemoryChunkInfo array `json:"MemoryChunkInfo,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The set of volumes and access capabilities specified for this connection.
	VolumeInfo array `json:"VolumeInfo,omitempty"`

}
