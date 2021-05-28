/* -----------------------------------------------------------------
 * enum_log_diagnostic_data_types.go -
 *
 * DMTF Redfish LogDiagnosticDataTypes enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type LogDiagnosticDataTypes string

const (
	// Manager diagnostic data.
	LogDiagnosticDataTypes_MANAGER LogDiagnosticDataTypes = "Manager"

	// Pre-OS diagnostic data.
	LogDiagnosticDataTypes_PRE_OS LogDiagnosticDataTypes = "PreOS"

	// Operating system (OS) diagnostic data.
	LogDiagnosticDataTypes_OS LogDiagnosticDataTypes = "OS"

	// OEM diagnostic data.
	LogDiagnosticDataTypes_OEM LogDiagnosticDataTypes = "OEM"

)
