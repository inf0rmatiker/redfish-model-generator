/* -----------------------------------------------------------------
* sensor_current_excerpt.go -
*
* DMTF Redfish SensorCurrentExcerpt resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Sensor schema describes a sensor and its properties.
type SensorCurrentExcerpt struct {
	// The crest factor for this sensor.
	CrestFactor float64 `json:"CrestFactor,omitempty"`

	// The link to the resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The total harmonic distortion (THD).
	THDPercent float64 `json:"THDPercent,omitempty"`

}
