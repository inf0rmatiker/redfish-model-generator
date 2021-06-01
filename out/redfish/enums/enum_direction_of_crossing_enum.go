/* -----------------------------------------------------------------
* enum_direction_of_crossing_enum.go -
*
* DMTF Redfish DirectionOfCrossingEnum enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type DirectionOfCrossingEnum string

const (
	// A trigger occurs when the metric value crosses the trigger value while increasing.
	DirectionOfCrossingEnum_INCREASING DirectionOfCrossingEnum = "Increasing"

	// A trigger occurs when the metric value crosses the trigger value while decreasing.
	DirectionOfCrossingEnum_DECREASING DirectionOfCrossingEnum = "Decreasing"

)
