/* -----------------------------------------------------------------
* sensor_energyk_wh_excerpt.go -
*
* DMTF Redfish SensorEnergykWhExcerpt resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Sensor schema describes a sensor and its properties.
type SensorEnergykWhExcerpt struct {
	// The link to the resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The total accumulation value for this sensor.
	LifetimeReading float64 `json:"LifetimeReading,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The date and time when the time-based properties were last reset.
	SensorResetTime string `json:"SensorResetTime,omitempty"`

}
