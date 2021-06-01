/* -----------------------------------------------------------------
* enum_reading_units.go -
*
* DMTF Redfish ReadingUnits enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReadingUnits string

const (
	// Indicates that the fan reading and thresholds are measured in rotations per minute.
	ReadingUnits_RPM ReadingUnits = "RPM"

	// Indicates that the fan reading and thresholds are measured in percentage.
	ReadingUnits_PERCENT ReadingUnits = "Percent"

)
