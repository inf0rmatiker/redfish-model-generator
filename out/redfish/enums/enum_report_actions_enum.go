/* -----------------------------------------------------------------
* enum_report_actions_enum.go -
*
* DMTF Redfish ReportActionsEnum enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReportActionsEnum string

const (
	// Record the occurrence to the metric report collection.
	ReportActionsEnum_LOG_TO_METRIC_REPORTS_COLLECTION ReportActionsEnum = "LogToMetricReportsCollection"

	// Send a Redfish event message containing the metric report.
	ReportActionsEnum_REDFISH_EVENT ReportActionsEnum = "RedfishEvent"

)
