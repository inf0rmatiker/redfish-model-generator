/* -----------------------------------------------------------------
* storage.go -
*
* DMTF Redfish Storage resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Storage - This schema defines a storage subsystem and its respective properties.  A storage subsystem represents a set of storage controllers (physical or virtual) and the resources such as volumes that can be accessed from that subsystem.
// Modeled after DMTF Redfish schema Storage
type Storage struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The set of drives attached to the storage controllers represented by this resource.
	Drives []Drive `json:"Drives,omitempty"`

	DrivesOdataCount int `json:"Drives@odata.count,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Redundancy information for the storage subsystem.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The set of storage controllers represented by this resource.
	StorageControllers []StorageController `json:"StorageControllers,omitempty"`

	StorageControllersOdataCount int `json:"StorageControllers@odata.count,omitempty"`

	// The set of volumes produced by the storage controllers represented by this resource.
	Volumes map[string]interface{} `json:"Volumes,omitempty"`

}
