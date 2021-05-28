/* -----------------------------------------------------------------
 * enum_circuit_type.go -
 *
 * DMTF Redfish CircuitType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type CircuitType string

const (
	// A mains input or utility circuit.
	CircuitType_MAINS CircuitType = "Mains"

	// A branch (output) circuit.
	CircuitType_BRANCH CircuitType = "Branch"

	// A subfeed (output) circuit.
	CircuitType_SUBFEED CircuitType = "Subfeed"

	// A feeder (output) circuit.
	CircuitType_FEEDER CircuitType = "Feeder"

)
