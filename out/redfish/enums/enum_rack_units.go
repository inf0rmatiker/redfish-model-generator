/* -----------------------------------------------------------------
* enum_rack_units.go -
*
* DMTF Redfish RackUnits enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type RackUnits string

const (
	// Defines a rack unit as being equal to 48 mm (1.89 in).
	RackUnits_OPEN_U RackUnits = "OpenU"

	// Defines a rack unit as being equal to 1.75 in (44.45 mm).
	RackUnits_EIA_310 RackUnits = "EIA_310"

)
