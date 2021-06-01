/* -----------------------------------------------------------------
* software_inventory.go -
*
* DMTF Redfish SoftwareInventory resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SoftwareInventory - This schema defines an inventory of software components.
// Modeled after DMTF Redfish schema SoftwareInventory
type SoftwareInventory struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// A string representing the lowest supported version of this software.
	LowestSupportedVersion string `json:"LowestSupportedVersion,omitempty"`

	// A string representing the manufacturer/producer of this software.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The ID(s) of the resources associated with this software inventory item.
	RelatedItem []idRef `json:"RelatedItem,omitempty"`

	RelatedItemOdataCount int `json:"RelatedItem@odata.count,omitempty"`

	// Release date of this software.
	ReleaseDate string `json:"ReleaseDate,omitempty"`

	// A string representing the implementation-specific ID for identifying this software.
	SoftwareId string `json:"SoftwareId,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// A list of strings representing the UEFI Device Path(s) of the component(s) associated with this software inventory item.
	UefiDevicePaths []string `json:"UefiDevicePaths,omitempty"`

	// Indicates whether this software can be updated by the update service.
	Updateable bool `json:"Updateable,omitempty"`

	// A string representing the version of this software.
	Version string `json:"Version,omitempty"`

}
