/* -----------------------------------------------------------------
 * enum_parity.go -
 *
 * DMTF Redfish Parity enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type Parity string

const (
	// No parity bit.
	Parity_NONE Parity = "None"

	// An even parity bit.
	Parity_EVEN Parity = "Even"

	// An odd parity bit.
	Parity_ODD Parity = "Odd"

	// A mark parity bit.
	Parity_MARK Parity = "Mark"

	// A space parity bit.
	Parity_SPACE Parity = "Space"

)
