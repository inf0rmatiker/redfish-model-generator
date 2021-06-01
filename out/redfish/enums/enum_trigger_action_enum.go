/* -----------------------------------------------------------------
* enum_trigger_action_enum.go -
*
* DMTF Redfish TriggerActionEnum enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type TriggerActionEnum string

const (
	// Upon a trigger, record in a log.
	TriggerActionEnum_LOG_TO_LOG_SERVICE TriggerActionEnum = "LogToLogService"

	// Upon a trigger, send a Redfish Event message of type Event.
	TriggerActionEnum_REDFISH_EVENT TriggerActionEnum = "RedfishEvent"

)
