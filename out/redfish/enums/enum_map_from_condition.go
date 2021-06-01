/* -----------------------------------------------------------------
* enum_map_from_condition.go -
*
* DMTF Redfish MapFromCondition enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type MapFromCondition string

const (
	// The logical operation for 'Equal'.
	MapFromCondition_EQU MapFromCondition = "EQU"

	// The logical operation for 'Not Equal'.
	MapFromCondition_NEQ MapFromCondition = "NEQ"

	// The logical operation for 'Greater than'.
	MapFromCondition_GTR MapFromCondition = "GTR"

	// The logical operation for 'Greater than or Equal'.
	MapFromCondition_GEQ MapFromCondition = "GEQ"

	// The logical operation for 'Less than'.
	MapFromCondition_LSS MapFromCondition = "LSS"

	// The logical operation for 'Less than or Equal'.
	MapFromCondition_LEQ MapFromCondition = "LEQ"

)
