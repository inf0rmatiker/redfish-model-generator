/* -----------------------------------------------------------------
* part_location.go -
*
* DMTF Redfish PartLocation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PartLocation - The part location within the placement.
// Modeled after DMTF Redfish schema PartLocation
type PartLocation struct {
	// The number that represents the location of the part.  If LocationType is slot and this unit is in slot 2 then the LocationOrdinalValue will be 2.
	LocationOrdinalValue float64 `json:"LocationOrdinalValue,omitempty"`

	// The type of location of the part, such as slot, bay, socket and slot.
	LocationType LocationType `json:"LocationType,omitempty"`

	// The orientation for the ordering of the slot enumeration used by the LocationOrdinalValue property.
	Orientation Orientation `json:"Orientation,omitempty"`

	// The reference point for the part location.  This is used to give guidance as to the general location of the part.
	Reference Reference `json:"Reference,omitempty"`

	// This is the label of the part location, such as a silk screened name or a printed label.
	ServiceLabel string `json:"ServiceLabel,omitempty"`

}
