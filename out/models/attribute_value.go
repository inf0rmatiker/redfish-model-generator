/* -----------------------------------------------------------------
* attribute_value.go -
*
* DMTF Redfish AttributeValue resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A possible value for an enumeration attribute.
type AttributeValue struct {
	// A user-readable display string of the value for the attribute in the defined language.
	ValueDisplayName string `json:"ValueDisplayName,omitempty"`

	// The unique value name for the attribute.
	ValueName string `json:"ValueName"`

}
