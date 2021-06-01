/* -----------------------------------------------------------------
* sensor.go -
*
* DMTF Redfish Sensor resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Sensor - The Sensor schema describes a sensor and its properties.
// Modeled after DMTF Redfish schema Sensor
type Sensor struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The estimated percent error of measured versus actual values.
	Accuracy float64 `json:"Accuracy,omitempty"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The adjusted maximum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMaxAllowableOperatingValue float64 `json:"AdjustedMaxAllowableOperatingValue,omitempty"`

	// The adjusted minimum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMinAllowableOperatingValue float64 `json:"AdjustedMinAllowableOperatingValue,omitempty"`

	// The product of voltage and current for an AC circuit, in Volt-Amperes units.
	ApparentVA float64 `json:"ApparentVA,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The combination of current-carrying conductors.
	ElectricalContext ElectricalContext `json:"ElectricalContext,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The power load utilization for this sensor.
	LoadPercent float64 `json:"LoadPercent,omitempty"`

	// The location information for this sensor.
	Location map[string]interface{} `json:"Location,omitempty"`

	// The maximum allowable operating value for this equipment.
	MaxAllowableOperatingValue float64 `json:"MaxAllowableOperatingValue,omitempty"`

	// The minimum allowable operating value for this equipment.
	MinAllowableOperatingValue float64 `json:"MinAllowableOperatingValue,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The peak sensor value.
	PeakReading float64 `json:"PeakReading,omitempty"`

	// The time when the peak sensor value occurred.
	PeakReadingTime string `json:"PeakReadingTime,omitempty"`

	// The area or device to which this sensor measurement applies.
	PhysicalContext map[string]interface{} `json:"PhysicalContext,omitempty"`

	// The usage or location within a device to which this sensor measurement applies.
	PhysicalSubContext map[string]interface{} `json:"PhysicalSubContext,omitempty"`

	// The power factor for this sensor.
	PowerFactor float64 `json:"PowerFactor,omitempty"`

	// The number of significant digits in the reading.
	Precision float64 `json:"Precision,omitempty"`

	// The square root of the difference term of squared ApparentVA and squared Power (Reading) for a circuit, in VAR units.
	ReactiveVAR float64 `json:"ReactiveVAR,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The maximum possible value for this sensor.
	ReadingRangeMax float64 `json:"ReadingRangeMax,omitempty"`

	// The minimum possible value for this sensor.
	ReadingRangeMin float64 `json:"ReadingRangeMin,omitempty"`

	// The type of sensor.
	ReadingType ReadingType `json:"ReadingType,omitempty"`

	// The units of the reading and thresholds.
	ReadingUnits string `json:"ReadingUnits,omitempty"`

	// The time interval between readings of the physical sensor.
	SensingFrequency float64 `json:"SensingFrequency,omitempty"`

	// The date and time when the time-based properties were last reset.
	SensorResetTime string `json:"SensorResetTime,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The set of thresholds defined for this sensor.
	Thresholds Thresholds `json:"Thresholds,omitempty"`

	// The voltage type for this sensor.
	VoltageType VoltageType `json:"VoltageType,omitempty"`

}
