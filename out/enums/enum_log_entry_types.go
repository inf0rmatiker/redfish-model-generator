/* -----------------------------------------------------------------
 * enum_log_entry_types.go -
 *
 * DMTF Redfish LogEntryTypes enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type LogEntryTypes string

const (
	// The log contains Redfish-defined messages.
	LogEntryTypes_EVENT LogEntryTypes = "Event"

	// The log contains legacy IPMI System Event Log (SEL) entries.
	LogEntryTypes_SEL LogEntryTypes = "SEL"

	// The log contains multiple log entry types and, therefore, the Log Service cannot guarantee a single entry type.
	LogEntryTypes_MULTIPLE LogEntryTypes = "Multiple"

	// The log contains entries in an OEM-defined format.
	LogEntryTypes_OEM LogEntryTypes = "OEM"

)
