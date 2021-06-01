/* -----------------------------------------------------------------
* power_supply.go -
*
* DMTF Redfish PowerSupply resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerSupply - Details of a power supplies associated with this system or device.
// Modeled after DMTF Redfish schema PowerSupply
type PowerSupply struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the identifier for the member within the collection.
	MemberId string `json:"MemberId,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// The Power Supply type (AC or DC).
	PowerSupplyType PowerSupplyType `json:"PowerSupplyType,omitempty"`

	// The line voltage type supported as an input to this Power Supply.
	LineInputVoltageType LineInputVoltageType `json:"LineInputVoltageType,omitempty"`

	// The line input voltage at which the Power Supply is operating.
	LineInputVoltage float64 `json:"LineInputVoltage,omitempty"`

	// The maximum capacity of this Power Supply.
	PowerCapacityWatts float64 `json:"PowerCapacityWatts,omitempty"`

	// The average power output of this Power Supply.
	LastPowerOutputWatts float64 `json:"LastPowerOutputWatts,omitempty"`

	// The model number for this Power Supply.
	Model string `json:"Model,omitempty"`

	// The firmware version for this Power Supply.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The serial number for this Power Supply.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The part number for this Power Supply.
	PartNumber string `json:"PartNumber,omitempty"`

	// The spare part number for this Power Supply.
	SparePartNumber string `json:"SparePartNumber,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

	RelatedItemOdataCount int `json:"RelatedItem@odata.count,omitempty"`

	RelatedItem@odata.navigationLink string `json:"RelatedItem@odata.navigationLink,omitempty"`

	// The ID(s) of the resources associated with this Power Limit.
	RelatedItem []idRef `json:"RelatedItem,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	Redundancy@odata.navigationLink string `json:"Redundancy@odata.navigationLink,omitempty"`

	// This structure is used to show redundancy for power supplies.  The Component ids will reference the members of the redundancy groups.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

	// This is the manufacturer of this power supply.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// This is the input ranges that the power supply can use.
	InputRanges []InputRange `json:"InputRanges,omitempty"`

}
