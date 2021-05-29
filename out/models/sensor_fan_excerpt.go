/* -----------------------------------------------------------------
* sensor_fan_excerpt.go -
*
* DMTF Redfish SensorFanExcerpt resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Sensor schema describes a sensor and its properties.
type SensorFanExcerpt struct {
	// The link to the resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The rotational speed.
	SpeedRPM float64 `json:"SpeedRPM,omitempty"`

}
