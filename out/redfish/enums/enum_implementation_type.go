/* -----------------------------------------------------------------
* enum_implementation_type.go -
*
* DMTF Redfish ImplementationType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ImplementationType string

const (
	// The metric is implemented as a physical sensor.
	ImplementationType_PHYSICAL_SENSOR ImplementationType = "PhysicalSensor"

	// The metric is implemented by applying a calculation on another metric property.  The calculation is specified in the CalculationAlgorithm property.
	ImplementationType_CALCULATED ImplementationType = "Calculated"

	// The metric is implemented by applying a calculation on one or more metric properties.  The calculation is not provided.
	ImplementationType_SYNTHESIZED ImplementationType = "Synthesized"

	// The metric is implemented as digital meter.
	ImplementationType_DIGITAL_METER ImplementationType = "DigitalMeter"

)
