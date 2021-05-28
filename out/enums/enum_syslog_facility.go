/* -----------------------------------------------------------------
 * enum_syslog_facility.go -
 *
 * DMTF Redfish SyslogFacility enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SyslogFacility string

const (
	// Kernel messages.
	SyslogFacility_KERN SyslogFacility = "Kern"

	// User-level messages.
	SyslogFacility_USER SyslogFacility = "User"

	// Mail system.
	SyslogFacility_MAIL SyslogFacility = "Mail"

	// System daemons.
	SyslogFacility_DAEMON SyslogFacility = "Daemon"

	// Security/authentication messages.
	SyslogFacility_AUTH SyslogFacility = "Auth"

	// Messages generated internally by syslogd.
	SyslogFacility_SYSLOG SyslogFacility = "Syslog"

	// Line printer subsystem.
	SyslogFacility_LPR SyslogFacility = "LPR"

	// Network news subsystem.
	SyslogFacility_NEWS SyslogFacility = "News"

	// UUCP subsystem.
	SyslogFacility_UUCP SyslogFacility = "UUCP"

	// Clock daemon.
	SyslogFacility_CRON SyslogFacility = "Cron"

	// Security/authentication messages.
	SyslogFacility_AUTHPRIV SyslogFacility = "Authpriv"

	// FTP daemon.
	SyslogFacility_FTP SyslogFacility = "FTP"

	// NTP subsystem.
	SyslogFacility_NTP SyslogFacility = "NTP"

	// Log audit.
	SyslogFacility_SECURITY SyslogFacility = "Security"

	// Log alert.
	SyslogFacility_CONSOLE SyslogFacility = "Console"

	// Scheduling daemon.
	SyslogFacility_SOLARIS_CRON SyslogFacility = "SolarisCron"

	// Locally used facility 0.
	SyslogFacility_LOCAL0 SyslogFacility = "Local0"

	// Locally used facility 1.
	SyslogFacility_LOCAL1 SyslogFacility = "Local1"

	// Locally used facility 2.
	SyslogFacility_LOCAL2 SyslogFacility = "Local2"

	// Locally used facility 3.
	SyslogFacility_LOCAL3 SyslogFacility = "Local3"

	// Locally used facility 4.
	SyslogFacility_LOCAL4 SyslogFacility = "Local4"

	// Locally used facility 5.
	SyslogFacility_LOCAL5 SyslogFacility = "Local5"

	// Locally used facility 6.
	SyslogFacility_LOCAL6 SyslogFacility = "Local6"

	// Locally used facility 7.
	SyslogFacility_LOCAL7 SyslogFacility = "Local7"

)
