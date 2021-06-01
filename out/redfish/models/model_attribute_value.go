/* -----------------------------------------------------------------
* attribute_value.go -
*
* DMTF Redfish AttributeValue resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// AttributeValue - A possible value for attributes of type 'Enumeration'.
// Modeled after DMTF Redfish schema AttributeValue
type AttributeValue struct {
	// A user-readable display string of the value of the attribute in the defined 'Language'.
	ValueDisplayName string `json:"ValueDisplayName,omitempty"`

	// The unique value name of the attribute.
	ValueName string `json:"ValueName"`

}
