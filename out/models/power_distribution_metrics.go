/* -----------------------------------------------------------------
* power_distribution_metrics.go -
*
* DMTF Redfish PowerDistributionMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This is the schema definition for the metrics of a power distribution component or unit, such as a floor power distribution unit (PDU) or switchgear.
type PowerDistributionMetrics struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Energy consumption (kWh).
	EnergykWh SensorEnergykWhExcerpt `json:"EnergykWh,omitempty"`

	// Humidity (percent).
	HumidityPercent SensorExcerpt `json:"HumidityPercent,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Power consumption (Watts).
	PowerWatts SensorPowerExcerpt `json:"PowerWatts,omitempty"`

	// Temperature (Celsius).
	TemperatureCelsius SensorExcerpt `json:"TemperatureCelsius,omitempty"`

}
