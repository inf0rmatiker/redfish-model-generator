/* -----------------------------------------------------------------
* storage_controller.go -
*
* DMTF Redfish StorageController resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// StorageController - This schema defines a storage controller and its respective properties.  A storage controller represents a storage device (physical or virtual) that produces Volumes.
// Modeled after DMTF Redfish schema StorageController
type StorageController struct {
	OdataId map[string]interface{} `json:"@odata.id"`

	// The user assigned asset tag for this storage controller.
	AssetTag string `json:"AssetTag,omitempty"`

	// The firmware version of this storage Controller.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The Durable names for the storage controller.
	Identifiers []Identifier `json:"Identifiers,omitempty"`

	// Contains references to other resources that are related to this resource.
	Links StorageControllerLinks `json:"Links,omitempty"`

	// This is the manufacturer of this storage controller.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// This is the identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// This is the model number for the storage controller.
	Model string `json:"Model,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this storage controller.
	PartNumber string `json:"PartNumber,omitempty"`

	// This is the SKU for this storage controller.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this storage controller.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The maximum speed of the storage controller's device interface.
	SpeedGbps float64 `json:"SpeedGbps,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// This represents the protocols by which this storage controller can be communicated to.
	SupportedControllerProtocols []Protocol `json:"SupportedControllerProtocols,omitempty"`

	// This represents the protocols which the storage controller can use to communicate with attached devices.
	SupportedDeviceProtocols []Protocol `json:"SupportedDeviceProtocols,omitempty"`

}
