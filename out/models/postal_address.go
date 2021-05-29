/* -----------------------------------------------------------------
* postal_address.go -
*
* DMTF Redfish PostalAddress resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The postal address for a resource.
type PostalAddress struct {
	// The additional code.
	AdditionalCode string `json:"AdditionalCode,omitempty"`

	// The room designation or other additional information.
	AdditionalInfo string `json:"AdditionalInfo,omitempty"`

	// The name of the building.
	Building string `json:"Building,omitempty"`

	// City, township, or shi (JP).
	City string `json:"City,omitempty"`

	// The postal community name.
	Community string `json:"Community,omitempty"`

	// The country.
	Country string `json:"Country,omitempty"`

	// A county, parish, gun (JP), or district (IN).
	District string `json:"District,omitempty"`

	// City division, borough, city district, ward, or chou (JP).
	Division string `json:"Division,omitempty"`

	// The floor.
	Floor string `json:"Floor,omitempty"`

	// The GPS coordinates of the part.
	GPSCoords string `json:"GPSCoords,omitempty"`

	// The numeric portion of house number.
	HouseNumber int `json:"HouseNumber,omitempty"`

	// The house number suffix.
	HouseNumberSuffix string `json:"HouseNumberSuffix,omitempty"`

	// The landmark.
	Landmark string `json:"Landmark,omitempty"`

	// A leading street direction.
	LeadingStreetDirection string `json:"LeadingStreetDirection,omitempty"`

	// The room designation or other additional information.
	Location string `json:"Location,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// Neighborhood or block.
	Neighborhood string `json:"Neighborhood,omitempty"`

	// The post office box (PO box).
	POBox string `json:"POBox,omitempty"`

	// The description of the type of place that is addressed.
	PlaceType string `json:"PlaceType,omitempty"`

	// The postal code or zip code.
	PostalCode string `json:"PostalCode,omitempty"`

	// The primary road or street.
	Road string `json:"Road,omitempty"`

	// The road branch.
	RoadBranch string `json:"RoadBranch,omitempty"`

	// The road post-modifier.
	RoadPostModifier string `json:"RoadPostModifier,omitempty"`

	// The road pre-modifier.
	RoadPreModifier string `json:"RoadPreModifier,omitempty"`

	// The road section.
	RoadSection string `json:"RoadSection,omitempty"`

	// The road sub branch.
	RoadSubBranch string `json:"RoadSubBranch,omitempty"`

	// The name or number of the room.
	Room string `json:"Room,omitempty"`

	// The seat, such as the desk, cubicle, or workstation.
	Seat string `json:"Seat,omitempty"`

	// Street name.
	Street string `json:"Street,omitempty"`

	// Avenue, Platz, Street, Circle.
	StreetSuffix string `json:"StreetSuffix,omitempty"`

	// A top-level subdivision within a country.
	Territory string `json:"Territory,omitempty"`

	// A trailing street suffix.
	TrailingStreetSuffix string `json:"TrailingStreetSuffix,omitempty"`

	// The name or number of the apartment unit or suite.
	Unit string `json:"Unit,omitempty"`

}
