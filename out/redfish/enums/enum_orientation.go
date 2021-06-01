/* -----------------------------------------------------------------
* enum_orientation.go -
*
* DMTF Redfish Orientation enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Orientation string

const (
	// Defines the ordering for the LocationOrdinalValue is front to back.
	Orientation_FRONT_TO_BACK Orientation = "FrontToBack"

	// Defines the ordering for the LocationOrdinalValue is back to front.
	Orientation_BACK_TO_FRONT Orientation = "BackToFront"

	// Defines the ordering for the LocationOrdinalValue is top to bottom.
	Orientation_TOP_TO_BOTTOM Orientation = "TopToBottom"

	// Defines the ordering for the LocationOrdinalValue is bottom to top.
	Orientation_BOTTOM_TO_TOP Orientation = "BottomToTop"

	// Defines the ordering for the LocationOrdinalValue is left to right.
	Orientation_LEFT_TO_RIGHT Orientation = "LeftToRight"

	// Defines the ordering for the LocationOrdinalValue is right to left.
	Orientation_RIGHT_TO_LEFT Orientation = "RightToLeft"

)
