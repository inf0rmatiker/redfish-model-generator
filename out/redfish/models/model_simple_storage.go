/* -----------------------------------------------------------------
* simple_storage.go -
*
* DMTF Redfish SimpleStorage resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SimpleStorage - This is the schema definition for the Simple Storage resource.  It represents the properties of a storage controller and its directly-attached devices.
// Modeled after DMTF Redfish schema SimpleStorage
type SimpleStorage struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The storage devices associated with this resource.
	Devices []Device `json:"Devices,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The UEFI device path used to access this storage controller.
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`

}
