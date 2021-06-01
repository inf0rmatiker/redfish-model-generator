/* -----------------------------------------------------------------
* pc_ie_device.go -
*
* DMTF Redfish PCIeDevice resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PCIeDevice - This is the schema definition for the PCIeDevice resource.  It represents the properties of a PCIeDevice attached to a System.
// Modeled after DMTF Redfish schema PCIeDevice
type PCIeDevice struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// A reference to the Assembly resource associated with this PCIe device.
	Assembly map[string]interface{} `json:"Assembly,omitempty"`

	// The user assigned asset tag for this PCIe device.
	AssetTag string `json:"AssetTag,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The device type for this PCIe device.
	DeviceType DeviceType `json:"DeviceType,omitempty"`

	// The version of firmware for this PCIe device.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links object contains the links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// This is the manufacturer of this PCIe device.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// This is the model number for the PCIe device.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this PCIe device.
	PartNumber string `json:"PartNumber,omitempty"`

	// This is the SKU for this PCIe device.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this PCIe device.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

}
