/* -----------------------------------------------------------------
* enum_expand.go -
*
* DMTF Redfish Expand enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Expand string

const (
	// Do not expand any references.
	Expand_NONE Expand = "None"

	// Expand all subordinate references.
	Expand_ALL Expand = "All"

	// Expand relevant subordinate references.  Relevant references are those that are tied to a constrained composition request, such as a request for a quantity of processors.
	Expand_RELEVANT Expand = "Relevant"

)
