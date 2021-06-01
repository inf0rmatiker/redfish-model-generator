/* -----------------------------------------------------------------
* enum_port_medium.go -
*
* DMTF Redfish PortMedium enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PortMedium string

const (
	// This port has an electrical cable connection.
	PortMedium_ELECTRICAL PortMedium = "Electrical"

	// This port has an optical cable connection.
	PortMedium_OPTICAL PortMedium = "Optical"

)
