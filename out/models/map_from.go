/* -----------------------------------------------------------------
* map_from.go -
*
* DMTF Redfish MapFrom resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A map-from condition for mapping dependency.
type MapFrom struct {
	// The attribute to use to evaluate this dependency expression.
	MapFromAttribute string `json:"MapFromAttribute,omitempty"`

	// The condition to use to evaluate this dependency expression.
	MapFromCondition MapFromCondition `json:"MapFromCondition,omitempty"`

	// The metadata property for the attribute that the MapFromAttribute property specifies to use to evaluate this dependency expression.
	MapFromProperty MapFromProperty `json:"MapFromProperty,omitempty"`

	// The value to use to evaluate this dependency expression.
	MapFromValue string `json:"MapFromValue,omitempty"`

	// The logical term that combines two or more map-from conditions in this dependency expression.  For example, `AND` for logical AND, or `OR` for logical OR.
	MapTerms MapTerms `json:"MapTerms,omitempty"`

}
