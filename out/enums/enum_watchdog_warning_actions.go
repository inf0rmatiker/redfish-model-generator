/* -----------------------------------------------------------------
 * enum_watchdog_warning_actions.go -
 *
 * DMTF Redfish WatchdogWarningActions enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type WatchdogWarningActions string

const (
	// No action taken.
	WatchdogWarningActions_NONE WatchdogWarningActions = "None"

	// Raise a (typically non-maskable) Diagnostic Interrupt.
	WatchdogWarningActions_DIAGNOSTIC_INTERRUPT WatchdogWarningActions = "DiagnosticInterrupt"

	// Raise a Systems Management Interrupt (SMI).
	WatchdogWarningActions_SMI WatchdogWarningActions = "SMI"

	// Raise a legacy IPMI messaging interrupt.
	WatchdogWarningActions_MESSAGING_INTERRUPT WatchdogWarningActions = "MessagingInterrupt"

	// Raise an interrupt using the ACPI System Control Interrupt (SCI).
	WatchdogWarningActions_SCI WatchdogWarningActions = "SCI"

	// Perform an OEM-defined action.
	WatchdogWarningActions_OEM WatchdogWarningActions = "OEM"

)
