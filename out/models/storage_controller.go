/* -----------------------------------------------------------------
* storage_controller.go -
*
* DMTF Redfish StorageController resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The StorageController schema describes a storage controller and its properties.  A storage controller represents a physical or virtual storage device that produces volumes.
type StorageController struct {
	OdataId string `json:"@odata.id"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly associated with this storage controller.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The user-assigned asset tag for this storage controller.
	AssetTag string `json:"AssetTag,omitempty"`

	// The cache memory of the storage controller in general detail.
	CacheSummary CacheSummary `json:"CacheSummary,omitempty"`

	// This property describes the various controller rates used for processes such as volume rebuild or consistency checks.
	ControllerRates Rates `json:"ControllerRates,omitempty"`

	// The firmware version of this storage controller.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The durable names for the storage controller.
	Identifiers array `json:"Identifiers,omitempty"`

	// The links to other resources that are related to this resource.
	Links StorageControllerLinks `json:"Links,omitempty"`

	// The location of the storage controller.
	Location Location `json:"Location,omitempty"`

	// The manufacturer of this storage controller.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// The model number for the storage controller.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The PCIe interface details for this controller.
	PCIeInterface PCIeInterface `json:"PCIeInterface,omitempty"`

	// The part number for this storage controller.
	PartNumber string `json:"PartNumber,omitempty"`

	// The link to the collection of ports that exist on the storage controller.
	Ports PortCollection `json:"Ports,omitempty"`

	// The SKU for this storage controller.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this storage controller.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The maximum speed of the storage controller's device interface.
	SpeedGbps float64 `json:"SpeedGbps,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The supported set of protocols for communicating to this storage controller.
	SupportedControllerProtocols array `json:"SupportedControllerProtocols,omitempty"`

	// The protocols that the storage controller can use to communicate with attached devices.
	SupportedDeviceProtocols array `json:"SupportedDeviceProtocols,omitempty"`

	// The set of RAID types supported by the storage controller.
	SupportedRAIDTypes array `json:"SupportedRAIDTypes,omitempty"`

}
