/* -----------------------------------------------------------------
* sensor_excerpt.go -
*
* DMTF Redfish SensorExcerpt resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SensorExcerpt - The Sensor schema describes a sensor and its properties.
// Modeled after DMTF Redfish schema SensorExcerpt
type SensorExcerpt struct {
	// The link to the Resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The peak sensor value.
	PeakReading float64 `json:"PeakReading,omitempty"`

	// The area or device to which this sensor measurement applies.
	PhysicalContext map[string]interface{} `json:"PhysicalContext,omitempty"`

	// The usage or location within a device to which this sensor measurement applies.
	PhysicalSubContext map[string]interface{} `json:"PhysicalSubContext,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The units of the reading and thresholds.
	ReadingUnits string `json:"ReadingUnits,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
