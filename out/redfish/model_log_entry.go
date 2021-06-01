/* -----------------------------------------------------------------
* log_entry.go -
*
* DMTF Redfish LogEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// LogEntry - The LogEntry schema defines the record format for a log.  It is designed for Redfish event logs, OEM-specific log formats, and the IPMI System Event Log (SEL).  The EntryType field indicates the type of log and the Resource includes several additional properties dependent on the EntryType.
// Modeled after DMTF Redfish schema LogEntry
type LogEntry struct {
}
