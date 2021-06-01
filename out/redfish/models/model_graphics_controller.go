/* -----------------------------------------------------------------
* graphics_controller.go -
*
* DMTF Redfish GraphicsController resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// GraphicsController - The GraphicsController schema defines a graphics controller that can be used to drive one or more display devices.
// Modeled after DMTF Redfish schema GraphicsController
type GraphicsController struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The user-assigned asset tag for this graphics controller.
	AssetTag string `json:"AssetTag,omitempty"`

	// The version of the graphics controller BIOS or primary graphics controller firmware.
	BiosVersion string `json:"BiosVersion,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The version of the graphics controller driver loaded in the operating system.
	DriverVersion string `json:"DriverVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the graphics controller.
	Location map[string]interface{} `json:"Location,omitempty"`

	// The manufacturer of this graphics controller.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The product model number of this graphics controller.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this graphics controller.
	PartNumber string `json:"PartNumber,omitempty"`

	// The ports of the graphics controller.
	Ports map[string]interface{} `json:"Ports,omitempty"`

	// The SKU for this graphics controller.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this graphics controller.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The spare part number of the graphics controller.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
