/* -----------------------------------------------------------------
* power_supply_metrics.go -
*
* DMTF Redfish PowerSupplyMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerSupplyMetrics - The PowerSupplyMetrics schema contains definitions for the metrics of a power supply.
// Modeled after DMTF Redfish schema PowerSupplyMetrics
type PowerSupplyMetrics struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The energy consumption of this unit.
	EnergykWh map[string]interface{} `json:"EnergykWh,omitempty"`

	// The fan speed reading for this power supply.
	FanSpeedPercent map[string]interface{} `json:"FanSpeedPercent,omitempty"`

	// The frequency reading for this power supply.
	FrequencyHz map[string]interface{} `json:"FrequencyHz,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The input current reading for this power supply.
	InputCurrentAmps map[string]interface{} `json:"InputCurrentAmps,omitempty"`

	// The input power reading for this power supply.
	InputPowerWatts map[string]interface{} `json:"InputPowerWatts,omitempty"`

	// The input voltage reading for this power supply.
	InputVoltage map[string]interface{} `json:"InputVoltage,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The total power output reading for this power supply.
	OutputPowerWatts map[string]interface{} `json:"OutputPowerWatts,omitempty"`

	// The current readings for this power supply.
	RailCurrentAmps []SensorCurrentExcerpt `json:"RailCurrentAmps,omitempty"`

	RailCurrentAmpsOdataCount int `json:"RailCurrentAmps@odata.count,omitempty"`

	// The power readings for this power supply.
	RailPowerWatts []SensorPowerExcerpt `json:"RailPowerWatts,omitempty"`

	RailPowerWattsOdataCount int `json:"RailPowerWatts@odata.count,omitempty"`

	// The voltage readings for this power supply.
	RailVoltage []SensorVoltageExcerpt `json:"RailVoltage,omitempty"`

	RailVoltageOdataCount int `json:"RailVoltage@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The temperature reading for this power supply.
	TemperatureCelsius map[string]interface{} `json:"TemperatureCelsius,omitempty"`

}
