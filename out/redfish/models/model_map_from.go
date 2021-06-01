/* -----------------------------------------------------------------
* map_from.go -
*
* DMTF Redfish MapFrom resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MapFrom - A map-from condition for mapping dependency.
// Modeled after DMTF Redfish schema MapFrom
type MapFrom struct {
	// The attribute that is used to evaluate this dependency expression.
	MapFromAttribute string `json:"MapFromAttribute,omitempty"`

	// The condition that is used to evaluate this dependency expression.
	MapFromCondition MapFromCondition `json:"MapFromCondition,omitempty"`

	// The meta-data property of the attribute specified in MapFromAttribute that is used to evaluate this dependency expression.
	MapFromProperty MapFromProperty `json:"MapFromProperty,omitempty"`

	// The value that the is used property specified in MapFromProperty that is used to evaluate this dependency expression.
	MapFromValue float64 `json:"MapFromValue,omitempty"`

	// The logical term used to combine two or more MapFrom conditions in this dependency expression.
	MapTerms MapTerms `json:"MapTerms,omitempty"`

}
