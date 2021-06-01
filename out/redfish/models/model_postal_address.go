/* -----------------------------------------------------------------
* postal_address.go -
*
* DMTF Redfish PostalAddress resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PostalAddress - The PostalAddress for a resource.
// Modeled after DMTF Redfish schema PostalAddress
type PostalAddress struct {
	// Additional code.
	AdditionalCode string `json:"AdditionalCode,omitempty"`

	// Name of the building.
	Building string `json:"Building,omitempty"`

	// City, township, or shi (JP).
	City string `json:"City,omitempty"`

	// Postal community name.
	Community string `json:"Community,omitempty"`

	// Country.
	Country string `json:"Country,omitempty"`

	// A county, parish, gun (JP), or  district (IN).
	District string `json:"District,omitempty"`

	// City division, borough, dity district, ward, chou (JP).
	Division string `json:"Division,omitempty"`

	// Floor.
	Floor string `json:"Floor,omitempty"`

	// The GPS coordinates of the part.
	GPSCoords string `json:"GPSCoords,omitempty"`

	// Numeric portion of house number.
	HouseNumber float64 `json:"HouseNumber,omitempty"`

	// House number suffix.
	HouseNumberSuffix string `json:"HouseNumberSuffix,omitempty"`

	// Landmark.
	Landmark string `json:"Landmark,omitempty"`

	// A leading street direction.
	LeadingStreetDirection string `json:"LeadingStreetDirection,omitempty"`

	// Room designation or other additional info.
	Location string `json:"Location,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// Neighborhood or block.
	Neighborhood string `json:"Neighborhood,omitempty"`

	// Post office box (P.O. box).
	POBox string `json:"POBox,omitempty"`

	// A description of the type of place that is addressed.
	PlaceType string `json:"PlaceType,omitempty"`

	// Postal code (or zip code).
	PostalCode string `json:"PostalCode,omitempty"`

	// A primary road or street.
	Road string `json:"Road,omitempty"`

	// Road branch.
	RoadBranch string `json:"RoadBranch,omitempty"`

	// Road post-modifier.
	RoadPostModifier string `json:"RoadPostModifier,omitempty"`

	// Road pre-modifier.
	RoadPreModifier string `json:"RoadPreModifier,omitempty"`

	// Road Section.
	RoadSection string `json:"RoadSection,omitempty"`

	// Road sub branch.
	RoadSubBranch string `json:"RoadSubBranch,omitempty"`

	// Name or number of the room.
	Room string `json:"Room,omitempty"`

	// Seat (desk, cubicle, workstation).
	Seat string `json:"Seat,omitempty"`

	// Street name.
	Street string `json:"Street,omitempty"`

	// Avenue, Platz, Street, Circle.
	StreetSuffix string `json:"StreetSuffix,omitempty"`

	// A top-level subdivision within a country.
	Territory string `json:"Territory,omitempty"`

	// A trailing street suffix.
	TrailingStreetSuffix string `json:"TrailingStreetSuffix,omitempty"`

	// Name or number of the unit (apartment, suite).
	Unit string `json:"Unit,omitempty"`

}
