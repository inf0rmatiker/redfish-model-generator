/* -----------------------------------------------------------------
* sensor_power_excerpt.go -
*
* DMTF Redfish SensorPowerExcerpt resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SensorPowerExcerpt - The Sensor schema describes a sensor and its properties.
// Modeled after DMTF Redfish schema SensorPowerExcerpt
type SensorPowerExcerpt struct {
	// The product of voltage and current for an AC circuit, in Volt-Amperes units.
	ApparentVA float64 `json:"ApparentVA,omitempty"`

	// The link to the Resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The power load utilization for this sensor.
	LoadPercent float64 `json:"LoadPercent,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The peak sensor value.
	PeakReading float64 `json:"PeakReading,omitempty"`

	// The area or device to which this sensor measurement applies.
	PhysicalContext map[string]interface{} `json:"PhysicalContext,omitempty"`

	// The usage or location within a device to which this sensor measurement applies.
	PhysicalSubContext map[string]interface{} `json:"PhysicalSubContext,omitempty"`

	// The power factor for this sensor.
	PowerFactor float64 `json:"PowerFactor,omitempty"`

	// The square root of the difference term of squared ApparentVA and squared Power (Reading) for a circuit, in VAR units.
	ReactiveVAR float64 `json:"ReactiveVAR,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The units of the reading and thresholds.
	ReadingUnits string `json:"ReadingUnits,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
