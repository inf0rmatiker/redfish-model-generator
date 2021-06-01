/* -----------------------------------------------------------------
* thermal_metrics.go -
*
* DMTF Redfish ThermalMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ThermalMetrics - The ThermalMetrics schema represents the thermal metrics of a chassis.
// Modeled after DMTF Redfish schema ThermalMetrics
type ThermalMetrics struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The temperature readings from all related sensors for this device.
	TemperatureReadingsCelsius []SensorArrayExcerpt `json:"TemperatureReadingsCelsius,omitempty"`

	TemperatureReadingsCelsiusOdataCount int `json:"TemperatureReadingsCelsius@odata.count,omitempty"`

	// The summary temperature readings for this chassis.
	TemperatureSummaryCelsius TemperatureSummary `json:"TemperatureSummaryCelsius,omitempty"`

}
