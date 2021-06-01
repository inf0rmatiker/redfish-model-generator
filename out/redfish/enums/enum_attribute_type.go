/* -----------------------------------------------------------------
* enum_attribute_type.go -
*
* DMTF Redfish AttributeType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AttributeType string

const (
	// The attributes that include a list of the known possible enumerated values.
	AttributeType_ENUMERATION AttributeType = "Enumeration"

	// The attributes that include free form text in their values.
	AttributeType_STRING AttributeType = "String"

	// The attributes that have integer numeric values.
	AttributeType_INTEGER AttributeType = "Integer"

	// The attributes that are true or false.
	AttributeType_BOOLEAN AttributeType = "Boolean"

	// The attributes that include password values and are not displayed as plain text. The value shall be null for GET requests.
	AttributeType_PASSWORD AttributeType = "Password"

)
