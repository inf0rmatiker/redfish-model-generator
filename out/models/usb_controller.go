/* -----------------------------------------------------------------
* usb_controller.go -
*
* DMTF Redfish USBController resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The USBController schema defines a Universal Serial Bus controller.
type USBController struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The manufacturer of this USB controller.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The product model number of this USB controller.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this USB controller.
	PartNumber string `json:"PartNumber,omitempty"`

	// The ports of the USB controller.
	Ports PortCollection `json:"Ports,omitempty"`

	// The SKU for this USB controller.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this USB controller.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The spare part number of the USB controller.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
