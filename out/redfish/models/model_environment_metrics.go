/* -----------------------------------------------------------------
* environment_metrics.go -
*
* DMTF Redfish EnvironmentMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EnvironmentMetrics - The EnvironmentMetrics schema represents the environmental metrics of a device.
// Modeled after DMTF Redfish schema EnvironmentMetrics
type EnvironmentMetrics struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Energy consumption (kWh).
	EnergykWh map[string]interface{} `json:"EnergykWh,omitempty"`

	// Fan speeds (percent).
	FanSpeedsPercent []SensorFanArrayExcerpt `json:"FanSpeedsPercent,omitempty"`

	FanSpeedsPercentOdataCount int `json:"FanSpeedsPercent@odata.count,omitempty"`

	// Humidity (percent).
	HumidityPercent map[string]interface{} `json:"HumidityPercent,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Power consumption (Watts).
	PowerWatts map[string]interface{} `json:"PowerWatts,omitempty"`

	// Temperature (Celsius).
	TemperatureCelsius map[string]interface{} `json:"TemperatureCelsius,omitempty"`

}
