/* -----------------------------------------------------------------
 * enum_voltage_type.go -
 *
 * DMTF Redfish VoltageType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type VoltageType string

const (
	// Alternating Current (AC) circuit.
	VoltageType_AC VoltageType = "AC"

	// Direct Current (DC) circuit.
	VoltageType_DC VoltageType = "DC"

)
