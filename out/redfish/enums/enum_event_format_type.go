/* -----------------------------------------------------------------
* enum_event_format_type.go -
*
* DMTF Redfish EventFormatType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EventFormatType string

const (
	// The subscription destination receives an event payload.
	EventFormatType_EVENT EventFormatType = "Event"

	// The subscription destination receives a metric report.
	EventFormatType_METRIC_REPORT EventFormatType = "MetricReport"

)
