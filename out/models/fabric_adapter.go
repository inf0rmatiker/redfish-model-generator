/* -----------------------------------------------------------------
* fabric_adapter.go -
*
* DMTF Redfish FabricAdapter resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A FabricAdapter represents the physical fabric adapter capable of connecting to an interconnect fabric.  Examples include but are not limited to Ethernet, NVMe over Fabrics, Gen-Z, and SAS fabric adapters.
type FabricAdapter struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The manufacturer name for the ASIC of this fabric adapter.
	ASICManufacturer string `json:"ASICManufacturer,omitempty"`

	// The part number for the ASIC on this fabric adapter.
	ASICPartNumber string `json:"ASICPartNumber,omitempty"`

	// The revision identifier for the ASIC on this fabric adapter.
	ASICRevisionIdentifier string `json:"ASICRevisionIdentifier,omitempty"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The firmware version of this fabric adapter.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The Gen-Z specific properties for this fabric adapter.
	GenZ GenZ `json:"GenZ,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links,omitempty"`

	// The manufacturer or OEM of this fabric adapter.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model string for this fabric adapter.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The PCIe interface details for this fabric adapter.
	PCIeInterface PCIeInterface `json:"PCIeInterface,omitempty"`

	// The part number for this fabric adapter.
	PartNumber string `json:"PartNumber,omitempty"`

	// The link to the collection of ports that exist on the fabric adapter.
	Ports PortCollection `json:"Ports,omitempty"`

	// The manufacturer SKU for this fabric adapter.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this fabric adapter.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The spare part number for this fabric adapter.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

	// The UUID for this fabric adapter.
	UUID UUID `json:"UUID,omitempty"`

}
