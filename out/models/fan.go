/* -----------------------------------------------------------------
* fan.go -
*
* DMTF Redfish Fan resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Fan schema describes a cooling fan unit for a computer system or similar devices contained within a chassis.
type Fan struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly associated with this fan.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An indication of whether this device can be inserted or removed while the equipment is in operation.
	HotPluggable bool `json:"HotPluggable,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The location of the fan.
	Location Location `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The manufacturer of this fan.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model number for this fan.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this fan.
	PartNumber string `json:"PartNumber,omitempty"`

	// The area or device associated with this fan.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

	// Power consumption (Watts).
	PowerWatts SensorPowerExcerpt `json:"PowerWatts,omitempty"`

	// The serial number for this fan.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The spare part number for this fan.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The fan speed reading.
	SpeedPercent SensorFanExcerpt `json:"SpeedPercent,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
