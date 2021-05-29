/* -----------------------------------------------------------------
* power_supply_metrics.go -
*
* DMTF Redfish PowerSupplyMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The PowerSupplyMetrics schema contains definitions for the metrics of a power supply.
type PowerSupplyMetrics struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The energy consumption of this unit.
	EnergykWh SensorEnergykWhExcerpt `json:"EnergykWh,omitempty"`

	// The fan speed reading for this power supply.
	FanSpeedPercent SensorFanExcerpt `json:"FanSpeedPercent,omitempty"`

	// The frequency reading for this power supply.
	FrequencyHz SensorExcerpt `json:"FrequencyHz,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The input current reading for this power supply.
	InputCurrentAmps SensorCurrentExcerpt `json:"InputCurrentAmps,omitempty"`

	// The input power reading for this power supply.
	InputPowerWatts SensorPowerExcerpt `json:"InputPowerWatts,omitempty"`

	// The input voltage reading for this power supply.
	InputVoltage SensorVoltageExcerpt `json:"InputVoltage,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The total power output reading for this power supply.
	OutputPowerWatts SensorPowerExcerpt `json:"OutputPowerWatts,omitempty"`

	// The current readings for this power supply.
	RailCurrentAmps array `json:"RailCurrentAmps,omitempty"`

	RailCurrentAmps@odata.count count `json:"RailCurrentAmps@odata.count,omitempty"`

	// The power readings for this power supply.
	RailPowerWatts array `json:"RailPowerWatts,omitempty"`

	RailPowerWatts@odata.count count `json:"RailPowerWatts@odata.count,omitempty"`

	// The voltage readings for this power supply.
	RailVoltage array `json:"RailVoltage,omitempty"`

	RailVoltage@odata.count count `json:"RailVoltage@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The temperature reading for this power supply.
	TemperatureCelsius SensorExcerpt `json:"TemperatureCelsius,omitempty"`

}
