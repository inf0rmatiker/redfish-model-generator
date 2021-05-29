/* -----------------------------------------------------------------
* placement.go -
*
* DMTF Redfish Placement resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The placement within the addressed location.
type Placement struct {
	// Area designation or other additional info.
	AdditionalInfo string `json:"AdditionalInfo,omitempty"`

	// The name of a rack location within a row.
	Rack string `json:"Rack,omitempty"`

	// The vertical location of the item, in terms of RackOffsetUnits.
	RackOffset int `json:"RackOffset,omitempty"`

	// The type of rack units in use.
	RackOffsetUnits RackUnits `json:"RackOffsetUnits,omitempty"`

	// The name of the row.
	Row string `json:"Row,omitempty"`

}
