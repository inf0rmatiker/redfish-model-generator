/* -----------------------------------------------------------------
* pc_ie_device.go -
*
* DMTF Redfish PCIeDevice resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The PCIeDevice schema describes the properties of a PCIe device that is attached to a system.
type PCIeDevice struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly associated with this PCIe device.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The user-assigned asset tag for this PCIe device.
	AssetTag string `json:"AssetTag,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The device type for this PCIe device.
	DeviceType DeviceType `json:"DeviceType,omitempty"`

	// The link to the environment metrics for this PCIe device.
	EnvironmentMetrics EnvironmentMetrics `json:"EnvironmentMetrics,omitempty"`

	// The version of firmware for this PCIe device.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The manufacturer of this PCIe device.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model number for the PCIe device.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to the collection of PCIe functions associated with this PCIe device.
	PCIeFunctions PCIeFunctionCollection `json:"PCIeFunctions,omitempty"`

	// The PCIe interface details for this PCIe device.
	PCIeInterface PCIeInterface `json:"PCIeInterface,omitempty"`

	// The part number for this PCIe device.
	PartNumber string `json:"PartNumber,omitempty"`

	// An indication of whether the PCIe device is prepared by the system for removal.
	ReadyToRemove bool `json:"ReadyToRemove,omitempty"`

	// The SKU for this PCIe device.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this PCIe device.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The spare part number of the PCIe device.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The UUID for this PCIe device.
	UUID UUID `json:"UUID,omitempty"`

}
