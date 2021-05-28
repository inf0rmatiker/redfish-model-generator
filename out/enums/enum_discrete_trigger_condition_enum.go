/* -----------------------------------------------------------------
 * enum_discrete_trigger_condition_enum.go -
 *
 * DMTF Redfish DiscreteTriggerConditionEnum enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DiscreteTriggerConditionEnum string

const (
	// A discrete trigger condition is met when the metric value becomes one of the values that the DiscreteTriggers property lists.
	DiscreteTriggerConditionEnum_SPECIFIED DiscreteTriggerConditionEnum = "Specified"

	// A discrete trigger condition is met whenever the metric value changes.
	DiscreteTriggerConditionEnum_CHANGED DiscreteTriggerConditionEnum = "Changed"

)
