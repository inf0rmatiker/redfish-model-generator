/* -----------------------------------------------------------------
* sensor.go -
*
* DMTF Redfish Sensor resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Sensor schema describes a sensor and its properties.
type Sensor struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The estimated percent error of measured versus actual values.
	Accuracy float64 `json:"Accuracy,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The adjusted maximum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMaxAllowableOperatingValue float64 `json:"AdjustedMaxAllowableOperatingValue,omitempty"`

	// The adjusted minimum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMinAllowableOperatingValue float64 `json:"AdjustedMinAllowableOperatingValue,omitempty"`

	// The product of voltage and current for an AC circuit, in Volt-Ampere units.
	ApparentVA float64 `json:"ApparentVA,omitempty"`

	// The crest factor for this sensor.
	CrestFactor float64 `json:"CrestFactor,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The combination of current-carrying conductors.
	ElectricalContext ElectricalContext `json:"ElectricalContext,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The implementation of the sensor.
	Implementation ImplementationType `json:"Implementation,omitempty"`

	// The total accumulation value for this sensor.
	LifetimeReading float64 `json:"LifetimeReading,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The power load utilization for this sensor.
	LoadPercent float64 `json:"LoadPercent,omitempty"`

	// The location information for this sensor.
	Location Location `json:"Location,omitempty"`

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
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

	// The usage or location within a device to which this sensor measurement applies.
	PhysicalSubContext PhysicalSubContext `json:"PhysicalSubContext,omitempty"`

	// The power factor for this sensor.
	PowerFactor float64 `json:"PowerFactor,omitempty"`

	// The number of significant digits in the reading.
	Precision float64 `json:"Precision,omitempty"`

	// The square root of the difference term of squared ApparentVA and squared Power (Reading) for a circuit, in var units.
	ReactiveVAR float64 `json:"ReactiveVAR,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The maximum possible value for this sensor.
	ReadingRangeMax float64 `json:"ReadingRangeMax,omitempty"`

	// The minimum possible value for this sensor.
	ReadingRangeMin float64 `json:"ReadingRangeMin,omitempty"`

	// The date and time that the reading was acquired from the sensor.
	ReadingTime string `json:"ReadingTime,omitempty"`

	// The type of sensor.
	ReadingType ReadingType `json:"ReadingType,omitempty"`

	// The units of the reading and thresholds.
	ReadingUnits string `json:"ReadingUnits,omitempty"`

	// An array of links to resources or objects that this sensor services.
	RelatedItem []map[string]string `json:"RelatedItem,omitempty"`

	RelatedItem@odata.count count `json:"RelatedItem@odata.count,omitempty"`

	// The time interval between readings of the physical sensor.
	SensingFrequency float64 `json:"SensingFrequency,omitempty"`

	// The time interval between readings of the sensor.
	SensingInterval string `json:"SensingInterval,omitempty"`

	// The date and time when the time-based properties were last reset.
	SensorResetTime string `json:"SensorResetTime,omitempty"`

	// The rotational speed.
	SpeedRPM float64 `json:"SpeedRPM,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The total harmonic distortion (THD).
	THDPercent float64 `json:"THDPercent,omitempty"`

	// The set of thresholds defined for this sensor.
	Thresholds Thresholds `json:"Thresholds,omitempty"`

	// The voltage type for this sensor.
	VoltageType VoltageType `json:"VoltageType,omitempty"`

}
