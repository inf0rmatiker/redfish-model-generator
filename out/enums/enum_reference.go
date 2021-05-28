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
	// The part is in the top of the unit.
	Reference_TOP Reference = "Top"

	// The part is in the bottom of the unit.
	Reference_BOTTOM Reference = "Bottom"

	// The part is in the front of the unit.
	Reference_FRONT Reference = "Front"

	// The part is in the rear of the unit.
	Reference_REAR Reference = "Rear"

	// The part is on the left side of of the unit.
	Reference_LEFT Reference = "Left"

	// The part is on the right side of the unit.
	Reference_RIGHT Reference = "Right"

	// The part is in the middle of the unit.
	Reference_MIDDLE Reference = "Middle"

)
