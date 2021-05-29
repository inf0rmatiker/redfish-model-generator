/* -----------------------------------------------------------------
* part_location.go -
*
* DMTF Redfish PartLocation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The part location for a resource within an enclosure.
type PartLocation struct {
	// The number that represents the location of the part.  For example, if LocationType is `Slot` and this unit is in slot 2, the LocationOrdinalValue is `2`.
	LocationOrdinalValue int `json:"LocationOrdinalValue,omitempty"`

	// The type of location of the part, such as slot, bay, socket, or slot.
	LocationType LocationType `json:"LocationType,omitempty"`

	// The orientation for the ordering of the slot enumeration used by the LocationOrdinalValue property.
	Orientation Orientation `json:"Orientation,omitempty"`

	// The reference point for the part location.  Provides guidance about the general location of the part.
	Reference Reference `json:"Reference,omitempty"`

	// The label of the part location, such as a silk-screened name or a printed label.
	ServiceLabel string `json:"ServiceLabel,omitempty"`

}
