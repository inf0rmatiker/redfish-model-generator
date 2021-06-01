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
	Parity_NONE Parity = "None"
	Parity_EVEN Parity = "Even"
	Parity_ODD Parity = "Odd"
	Parity_MARK Parity = "Mark"
	Parity_SPACE Parity = "Space"
)
