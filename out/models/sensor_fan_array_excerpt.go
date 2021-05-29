/* -----------------------------------------------------------------
* sensor_fan_array_excerpt.go -
*
* DMTF Redfish SensorFanArrayExcerpt resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Sensor schema describes a sensor and its properties.
type SensorFanArrayExcerpt struct {
	// The link to the resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The name of the device.
	DeviceName string `json:"DeviceName,omitempty"`

	// The area or device to which this sensor measurement applies.
	PhysicalContext PhysicalContext `json:"PhysicalContext,omitempty"`

	// The usage or location within a device to which this sensor measurement applies.
	PhysicalSubContext PhysicalSubContext `json:"PhysicalSubContext,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

	// The rotational speed.
	SpeedRPM float64 `json:"SpeedRPM,omitempty"`

}
