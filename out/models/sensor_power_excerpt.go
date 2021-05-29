/* -----------------------------------------------------------------
* sensor_power_excerpt.go -
*
* DMTF Redfish SensorPowerExcerpt resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Sensor schema describes a sensor and its properties.
type SensorPowerExcerpt struct {
	// The product of voltage and current for an AC circuit, in Volt-Ampere units.
	ApparentVA float64 `json:"ApparentVA,omitempty"`

	// The link to the resource that provides the data for this sensor.
	DataSourceUri string `json:"DataSourceUri,omitempty"`

	// The power factor for this sensor.
	PowerFactor float64 `json:"PowerFactor,omitempty"`

	// The square root of the difference term of squared ApparentVA and squared Power (Reading) for a circuit, in var units.
	ReactiveVAR float64 `json:"ReactiveVAR,omitempty"`

	// The sensor value.
	Reading float64 `json:"Reading,omitempty"`

}
