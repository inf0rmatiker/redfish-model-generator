/* -----------------------------------------------------------------
 * enum_metric_report_definition_type.go -
 *
 * DMTF Redfish MetricReportDefinitionType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MetricReportDefinitionType string

const (
	// The metric report is generated at a periodic time interval, specified in the Schedule property.
	MetricReportDefinitionType_PERIODIC MetricReportDefinitionType = "Periodic"

	// The metric report is generated when any of the metric values change.
	MetricReportDefinitionType_ON_CHANGE MetricReportDefinitionType = "OnChange"

	// The metric report is generated when a HTTP GET is performed on the specified metric report.
	MetricReportDefinitionType_ON_REQUEST MetricReportDefinitionType = "OnRequest"

)
