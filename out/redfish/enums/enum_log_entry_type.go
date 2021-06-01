/* -----------------------------------------------------------------
* enum_log_entry_type.go -
*
* DMTF Redfish LogEntryType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type LogEntryType string

const (
	// A Redfish-defined message.
	LogEntryType_EVENT LogEntryType = "Event"

	// A legacy IPMI System Event Log (SEL) entry.
	LogEntryType_SEL LogEntryType = "SEL"

	// An entry in an OEM-defined format.
	LogEntryType_OEM LogEntryType = "Oem"

)
