/* -----------------------------------------------------------------
* enum_health.go -
*
* DMTF Redfish Health enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Health string

const (
	// Normal.
	Health_OK Health = "OK"

	// A condition requires attention.
	Health_WARNING Health = "Warning"

	// A critical condition requires immediate attention.
	Health_CRITICAL Health = "Critical"

)
