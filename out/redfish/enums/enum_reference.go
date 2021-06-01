/* -----------------------------------------------------------------
* enum_reference.go -
*
* DMTF Redfish Reference enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Reference string

const (
	// Defines the part as being in the top of the unit.
	Reference_TOP Reference = "Top"

	// Defines the part as being in the bottom of the unit.
	Reference_BOTTOM Reference = "Bottom"

	// Defines the part as being in the front of the unit.
	Reference_FRONT Reference = "Front"

	// Defines the part as being in the rear of the unit.
	Reference_REAR Reference = "Rear"

	// Defines the part as being in the left of the unit.
	Reference_LEFT Reference = "Left"

	// Defines the part as being in the right of the unit.
	Reference_RIGHT Reference = "Right"

	// Defines the part as being in the middle of the unit.
	Reference_MIDDLE Reference = "Middle"

)
