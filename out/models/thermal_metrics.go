/* -----------------------------------------------------------------
* thermal_metrics.go -
*
* DMTF Redfish ThermalMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ThermalMetrics schema represents the thermal metrics of a chassis.
type ThermalMetrics struct {
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

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The temperature readings from all related sensors for this device.
	TemperatureReadingsCelsius array `json:"TemperatureReadingsCelsius,omitempty"`

	TemperatureReadingsCelsius@odata.count count `json:"TemperatureReadingsCelsius@odata.count,omitempty"`

	// The summary temperature readings for this chassis.
	TemperatureSummaryCelsius TemperatureSummary `json:"TemperatureSummaryCelsius,omitempty"`

}
