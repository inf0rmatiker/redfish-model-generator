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
	// A list of the known possible enumerated values.
	AttributeType_ENUMERATION AttributeType = "Enumeration"

	// Free-form text in their values.
	AttributeType_STRING AttributeType = "String"

	// An integer value.
	AttributeType_INTEGER AttributeType = "Integer"

	// A flag with a `true` or `false` value.
	AttributeType_BOOLEAN AttributeType = "Boolean"

	// Password values that do not appear as plain text.  The value shall be null in responses.
	AttributeType_PASSWORD AttributeType = "Password"

)
