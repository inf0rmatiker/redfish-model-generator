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
	// When a trigger condition is met, record in a log.
	TriggerActionEnum_LOG_TO_LOG_SERVICE TriggerActionEnum = "LogToLogService"

	// When a trigger condition is met, the service sends an event to subscribers.
	TriggerActionEnum_REDFISH_EVENT TriggerActionEnum = "RedfishEvent"

	// When a trigger condition is met, force an update of the specified metric reports.
	TriggerActionEnum_REDFISH_METRIC_REPORT TriggerActionEnum = "RedfishMetricReport"

)
