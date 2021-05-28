/* -----------------------------------------------------------------
 * enum_clearing_type.go -
 *
 * DMTF Redfish ClearingType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ClearingType string

const (
	// This enumeration shall describe when the message for an event is cleared by the other messages in the ClearingLogic property, provided the OriginOfCondition for both events are the same.
	ClearingType_SAME_ORIGIN_OF_CONDITION ClearingType = "SameOriginOfCondition"

)
