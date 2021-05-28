/* -----------------------------------------------------------------
 * enum_syslog_severity.go -
 *
 * DMTF Redfish SyslogSeverity enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SyslogSeverity string

const (
	// A panic condition.
	SyslogSeverity_EMERGENCY SyslogSeverity = "Emergency"

	// A condition that should be corrected immediately, such as a corrupted system database.
	SyslogSeverity_ALERT SyslogSeverity = "Alert"

	// Hard device errors.
	SyslogSeverity_CRITICAL SyslogSeverity = "Critical"

	// An Error.
	SyslogSeverity_ERROR SyslogSeverity = "Error"

	// A Warning.
	SyslogSeverity_WARNING SyslogSeverity = "Warning"

	// Conditions that are not error conditions, but that may require special handling.
	SyslogSeverity_NOTICE SyslogSeverity = "Notice"

	// Informational only.
	SyslogSeverity_INFORMATIONAL SyslogSeverity = "Informational"

	// Messages that contain information normally of use only when debugging a program.
	SyslogSeverity_DEBUG SyslogSeverity = "Debug"

	// A message of any severity.
	SyslogSeverity_ALL SyslogSeverity = "All"

)
