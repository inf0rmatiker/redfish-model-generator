/* -----------------------------------------------------------------
* nv_me_firmware_image.go -
*
* SNIA Swordfish NVMeFirmwareImage resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NVMeFirmwareImage - NVMe Domain firmware image information.
// Modeled after SNIA Swordfish schema NVMeFirmwareImage
type NVMeFirmwareImage struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The firmware version of the available NVMe firmware image.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The type of NVMe Device this image is associated with.
	NVMeDeviceType map[string]interface{} `json:"NVMeDeviceType,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The vendor or manufacturer associated with this NVMe firmware image.
	Vendor string `json:"Vendor,omitempty"`

}
