/* -----------------------------------------------------------------
* power_supply.go -
*
* DMTF Redfish PowerSupply resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The PowerSupply schema describes a power supply unit.
type PowerSupply struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly associated with this power supply.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The efficiency ratings of this power supply.
	EfficiencyRatings array `json:"EfficiencyRatings,omitempty"`

	// The firmware version for this power supply.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// An indication of whether this device can be inserted or removed while the equipment is in operation.
	HotPluggable bool `json:"HotPluggable,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The nominal voltage type of the line input to this power supply.
	InputNominalVoltageType NominalVoltageType `json:"InputNominalVoltageType,omitempty"`

	// The input ranges that the power supply can use.
	InputRanges array `json:"InputRanges,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the power supply.
	Location Location `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool `json:"LocationIndicatorActive,omitempty"`

	// The manufacturer of this power supply.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The link to the power supply metrics resource associated with this power supply.
	Metrics PowerSupplyMetrics `json:"Metrics,omitempty"`

	// The model number for this power supply.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The output power rails provided by this power supply.
	OutputRails array `json:"OutputRails,omitempty"`

	// The part number for this power supply.
	PartNumber string `json:"PartNumber,omitempty"`

	// The number of ungrounded current-carrying conductors (phases) and the total number of conductors (wires) provided for the power supply input connector.
	PhaseWiringType PhaseWiringType `json:"PhaseWiringType,omitempty"`

	// The type of plug according to NEMA, IEC, or regional standards.
	PlugType PlugType `json:"PlugType,omitempty"`

	// The maximum capacity of this power supply.
	PowerCapacityWatts float64 `json:"PowerCapacityWatts,omitempty"`

	// The power supply type (AC or DC).
	PowerSupplyType PowerSupplyType `json:"PowerSupplyType,omitempty"`

	// The production or manufacturing date of this power supply.
	ProductionDate string `json:"ProductionDate,omitempty"`

	// The serial number for this power supply.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The spare part number for this power supply.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The hardware version of this power supply.
	Version string `json:"Version,omitempty"`

}
