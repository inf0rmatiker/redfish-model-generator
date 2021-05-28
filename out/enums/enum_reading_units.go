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
	// The fan reading and thresholds are measured in revolutions per minute.
	ReadingUnits_RPM ReadingUnits = "RPM"

	// The fan reading and thresholds are measured as a percentage.
	ReadingUnits_PERCENT ReadingUnits = "Percent"

)
