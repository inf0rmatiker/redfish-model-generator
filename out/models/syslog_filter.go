/* -----------------------------------------------------------------
* syslog_filter.go -
*
* DMTF Redfish SyslogFilter resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A syslog filter.
type SyslogFilter struct {
	// The types of programs that can log messages.
	LogFacilities array `json:"LogFacilities,omitempty"`

	// The lowest severity level message that will be logged.
	LowestSeverity SyslogSeverity `json:"LowestSeverity,omitempty"`

}
