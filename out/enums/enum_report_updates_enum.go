/* -----------------------------------------------------------------
 * enum_report_updates_enum.go -
 *
 * DMTF Redfish ReportUpdatesEnum enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ReportUpdatesEnum string

const (
	// Overwrite the metric report.
	ReportUpdatesEnum_OVERWRITE ReportUpdatesEnum = "Overwrite"

	// New information is appended to the metric report.  The metric report entries are overwritten with new entries when the metric report has reached its maximum capacity.
	ReportUpdatesEnum_APPEND_WRAPS_WHEN_FULL ReportUpdatesEnum = "AppendWrapsWhenFull"

	// New information is appended to the metric report.  The service stops adding entries when the metric report has reached its maximum capacity.
	ReportUpdatesEnum_APPEND_STOPS_WHEN_FULL ReportUpdatesEnum = "AppendStopsWhenFull"

	// A new metric report is created, whose resource name is the metric report resource name concatenated with the timestamp.
	ReportUpdatesEnum_NEW_REPORT ReportUpdatesEnum = "NewReport"

)
