/* -----------------------------------------------------------------
 * enum_map_terms.go -
 *
 * DMTF Redfish MapTerms enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MapTerms string

const (
	// The operation used for logical 'AND' of dependency terms.
	MapTerms_AND MapTerms = "AND"

	// The operation used for logical 'OR' of dependency terms.
	MapTerms_OR MapTerms = "OR"

)
