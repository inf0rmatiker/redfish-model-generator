/* -----------------------------------------------------------------
* enum_log_entry_code.go -
*
* DMTF Redfish LogEntryCode enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type LogEntryCode string

const (
	// The condition has been asserted.
	LogEntryCode_ASSERT LogEntryCode = "Assert"

	// The condition has been deasserted.
	LogEntryCode_DEASSERT LogEntryCode = "Deassert"

	// The reading crossed the Lower Non-critical threshold while going low.
	LogEntryCode_LOWER _NON-CRITICAL - GOING LOW LogEntryCode = "Lower Non-critical - going low"

	// The reading crossed the Lower Non-critical threshold while going high.
	LogEntryCode_LOWER _NON-CRITICAL - GOING HIGH LogEntryCode = "Lower Non-critical - going high"

	// The reading crossed the Lower Critical threshold while going low.
	LogEntryCode_LOWER _CRITICAL - GOING LOW LogEntryCode = "Lower Critical - going low"

	// The reading crossed the Lower Critical threshold while going high.
	LogEntryCode_LOWER _CRITICAL - GOING HIGH LogEntryCode = "Lower Critical - going high"

	// The reading crossed the Lower Non-recoverable threshold while going low.
	LogEntryCode_LOWER _NON-RECOVERABLE - GOING LOW LogEntryCode = "Lower Non-recoverable - going low"

	// The reading crossed the Lower Non-recoverable threshold while going high.
	LogEntryCode_LOWER _NON-RECOVERABLE - GOING HIGH LogEntryCode = "Lower Non-recoverable - going high"

	// The reading crossed the Upper Non-critical threshold while going low.
	LogEntryCode_UPPER _NON-CRITICAL - GOING LOW LogEntryCode = "Upper Non-critical - going low"

	// The reading crossed the Upper Non-critical threshold while going high.
	LogEntryCode_UPPER _NON-CRITICAL - GOING HIGH LogEntryCode = "Upper Non-critical - going high"

	// The reading crossed the Upper Critical threshold while going low.
	LogEntryCode_UPPER _CRITICAL - GOING LOW LogEntryCode = "Upper Critical - going low"

	// The reading crossed the Upper Critical threshold while going high.
	LogEntryCode_UPPER _CRITICAL - GOING HIGH LogEntryCode = "Upper Critical - going high"

	// The reading crossed the Upper Non-recoverable threshold while going low.
	LogEntryCode_UPPER _NON-RECOVERABLE - GOING LOW LogEntryCode = "Upper Non-recoverable - going low"

	// The reading crossed the Upper Non-recoverable threshold while going high.
	LogEntryCode_UPPER _NON-RECOVERABLE - GOING HIGH LogEntryCode = "Upper Non-recoverable - going high"

	// The state transitioned to idle.
	LogEntryCode_TRANSITION TO _IDLE LogEntryCode = "Transition to Idle"

	// The state transitioned to active.
	LogEntryCode_TRANSITION TO _ACTIVE LogEntryCode = "Transition to Active"

	// The state transitioned to busy.
	LogEntryCode_TRANSITION TO _BUSY LogEntryCode = "Transition to Busy"

	// The state has been deasserted.
	LogEntryCode_STATE _DEASSERTED LogEntryCode = "State Deasserted"

	// The state has been asserted.
	LogEntryCode_STATE _ASSERTED LogEntryCode = "State Asserted"

	// A Predictive Failure is no longer present.
	LogEntryCode_PREDICTIVE _FAILURE DEASSERTED LogEntryCode = "Predictive Failure deasserted"

	// A Predictive Failure has been detected.
	LogEntryCode_PREDICTIVE _FAILURE ASSERTED LogEntryCode = "Predictive Failure asserted"

	// A limit has not been exceeded.
	LogEntryCode_LIMIT _NOT _EXCEEDED LogEntryCode = "Limit Not Exceeded"

	// A limit has been exceeded.
	LogEntryCode_LIMIT _EXCEEDED LogEntryCode = "Limit Exceeded"

	// Performance meets expectations.
	LogEntryCode_PERFORMANCE _MET LogEntryCode = "Performance Met"

	// Performance does not meet expectations.
	LogEntryCode_PERFORMANCE _LAGS LogEntryCode = "Performance Lags"

	// A state has changed to OK.
	LogEntryCode_TRANSITION TO OK LogEntryCode = "Transition to OK"

	// A state has changed to Non-Critical from OK.
	LogEntryCode_TRANSITION TO _NON-_CRITICAL FROM OK LogEntryCode = "Transition to Non-Critical from OK"

	// A state has changed to Critical from less severe.
	LogEntryCode_TRANSITION TO _CRITICAL FROM LESS SEVERE LogEntryCode = "Transition to Critical from less severe"

	// A state has changed to Non-recoverable from less severe.
	LogEntryCode_TRANSITION TO _NON-RECOVERABLE FROM LESS SEVERE LogEntryCode = "Transition to Non-recoverable from less severe"

	// A state has changed to Non-Critical from more severe.
	LogEntryCode_TRANSITION TO _NON-_CRITICAL FROM MORE SEVERE LogEntryCode = "Transition to Non-Critical from more severe"

	// A state has changed to Critical from Non-recoverable.
	LogEntryCode_TRANSITION TO _CRITICAL FROM _NON-RECOVERABLE LogEntryCode = "Transition to Critical from Non-recoverable"

	// A state has changed to Non-recoverable.
	LogEntryCode_TRANSITION TO _NON-RECOVERABLE LogEntryCode = "Transition to Non-recoverable"

	// A monitor event.
	LogEntryCode_MONITOR LogEntryCode = "Monitor"

	// An informational event.
	LogEntryCode_INFORMATIONAL LogEntryCode = "Informational"

	// A device has been removed or is absent.
	LogEntryCode_DEVICE _REMOVED / _DEVICE _ABSENT LogEntryCode = "Device Removed / Device Absent"

	// A device has been inserted or is present.
	LogEntryCode_DEVICE _INSERTED / _DEVICE _PRESENT LogEntryCode = "Device Inserted / Device Present"

	// A device has been disabled.
	LogEntryCode_DEVICE _DISABLED LogEntryCode = "Device Disabled"

	// A device has been enabled.
	LogEntryCode_DEVICE _ENABLED LogEntryCode = "Device Enabled"

	// A state has transitioned to Running.
	LogEntryCode_TRANSITION TO _RUNNING LogEntryCode = "Transition to Running"

	// A state has transitioned to In Test.
	LogEntryCode_TRANSITION TO _IN _TEST LogEntryCode = "Transition to In Test"

	// A state has transitioned to Power Off.
	LogEntryCode_TRANSITION TO _POWER _OFF LogEntryCode = "Transition to Power Off"

	// A state has transitioned to On Line.
	LogEntryCode_TRANSITION TO _ON _LINE LogEntryCode = "Transition to On Line"

	// A state has transitioned to Off Line.
	LogEntryCode_TRANSITION TO _OFF _LINE LogEntryCode = "Transition to Off Line"

	// A state has transitioned to Off Duty.
	LogEntryCode_TRANSITION TO _OFF _DUTY LogEntryCode = "Transition to Off Duty"

	// A state has transitioned to Degraded.
	LogEntryCode_TRANSITION TO _DEGRADED LogEntryCode = "Transition to Degraded"

	// A state has transitioned to Power Save.
	LogEntryCode_TRANSITION TO _POWER _SAVE LogEntryCode = "Transition to Power Save"

	// An install error has been detected.
	LogEntryCode_INSTALL _ERROR LogEntryCode = "Install Error"

	// Indicates that full redundancy has been regained.
	LogEntryCode_FULLY _REDUNDANT LogEntryCode = "Fully Redundant"

	// Entered any non-redundant state, including Non-redundant: Insufficient Resources.
	LogEntryCode_REDUNDANCY _LOST LogEntryCode = "Redundancy Lost"

	// Redundancy still exists, but at less than full level.
	LogEntryCode_REDUNDANCY _DEGRADED LogEntryCode = "Redundancy Degraded"

	// Redundancy has been lost but unit is functioning with minimum resources needed for normal operation.
	LogEntryCode_NON-REDUNDANT:_SUFFICIENT _RESOURCES FROM _REDUNDANT LogEntryCode = "Non-redundant:Sufficient Resources from Redundant"

	// Unit has regained minimum resources needed for normal operation.
	LogEntryCode_NON-REDUNDANT:_SUFFICIENT _RESOURCES FROM _INSUFFICIENT _RESOURCES LogEntryCode = "Non-redundant:Sufficient Resources from Insufficient Resources"

	// Unit is non-redundant and has insufficient resources to maintain normal operation.
	LogEntryCode_NON-REDUNDANT:_INSUFFICIENT _RESOURCES LogEntryCode = "Non-redundant:Insufficient Resources"

	// Unit has lost some redundant resource(s) but is still in a redundant state.
	LogEntryCode_REDUNDANCY _DEGRADED FROM _FULLY _REDUNDANT LogEntryCode = "Redundancy Degraded from Fully Redundant"

	// Unit has regained some resource(s) and is redundant but not fully redundant.
	LogEntryCode_REDUNDANCY _DEGRADED FROM _NON-REDUNDANT LogEntryCode = "Redundancy Degraded from Non-redundant"

	// The ACPI-defined D0 power state.
	LogEntryCode_D0 _POWER _STATE LogEntryCode = "D0 Power State"

	// The ACPI-defined D1 power state.
	LogEntryCode_D1 _POWER _STATE LogEntryCode = "D1 Power State"

	// The ACPI-defined D2 power state.
	LogEntryCode_D2 _POWER _STATE LogEntryCode = "D2 Power State"

	// The ACPI-defined D3 power state.
	LogEntryCode_D3 _POWER _STATE LogEntryCode = "D3 Power State"

)
