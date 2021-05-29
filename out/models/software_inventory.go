/* -----------------------------------------------------------------
* software_inventory.go -
*
* DMTF Redfish SoftwareInventory resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The SoftwareInventory schema contains an inventory of software components.  This can include software components such as BIOS, BMC firmware, firmware for other devices, system drivers, or provider software.
type SoftwareInventory struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The lowest supported version of this software.
	LowestSupportedVersion string `json:"LowestSupportedVersion,omitempty"`

	// The manufacturer or producer of this software.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// A DSP0274-defined measurement block.
	Measurement MeasurementBlock `json:"Measurement,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The IDs of the Resources associated with this software inventory item.
	RelatedItem []map[string]string `json:"RelatedItem,omitempty"`

	RelatedItem@odata.count count `json:"RelatedItem@odata.count,omitempty"`

	// The release date of this software.
	ReleaseDate string `json:"ReleaseDate,omitempty"`

	// The implementation-specific label that identifies this software.
	SoftwareId string `json:"SoftwareId,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

	// The list of UEFI device paths of the components associated with this software inventory item.
	UefiDevicePaths []string `json:"UefiDevicePaths,omitempty"`

	// An indication of whether the Update Service can update this software.
	Updateable bool `json:"Updateable,omitempty"`

	// The version of this software.
	Version string `json:"Version,omitempty"`

	// Indicates if the software is write-protected.
	WriteProtected bool `json:"WriteProtected,omitempty"`

}
