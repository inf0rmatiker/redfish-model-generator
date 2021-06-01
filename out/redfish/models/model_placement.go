/* -----------------------------------------------------------------
* placement.go -
*
* DMTF Redfish Placement resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Placement - The placement within the addressed location.
// Modeled after DMTF Redfish schema Placement
type Placement struct {
	// Name of a rack location within a row.
	Rack string `json:"Rack,omitempty"`

	// Vertical location of the item in terms of RackOffsetUnits.
	RackOffset float64 `json:"RackOffset,omitempty"`

	// The type of Rack Units in use.
	RackOffsetUnits RackUnits `json:"RackOffsetUnits,omitempty"`

	// Name of row.
	Row string `json:"Row,omitempty"`

}
