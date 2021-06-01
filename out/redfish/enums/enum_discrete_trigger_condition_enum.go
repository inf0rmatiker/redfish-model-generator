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
	// A discrete trigger occurs when the value of the metric becomes one of the values listed in the DiscreteTriggers property.
	DiscreteTriggerConditionEnum_SPECIFIED DiscreteTriggerConditionEnum = "Specified"

	// A discrete trigger occures whenever the value of the metric changes.
	DiscreteTriggerConditionEnum_CHANGED DiscreteTriggerConditionEnum = "Changed"

)
