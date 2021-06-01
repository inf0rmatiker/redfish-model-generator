/* -----------------------------------------------------------------
* enum_event_severity.go -
*
* DMTF Redfish EventSeverity enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EventSeverity string

const (
	// Informational or operating normally.
	EventSeverity_OK EventSeverity = "OK"

	// A condition that requires attention.
	EventSeverity_WARNING EventSeverity = "Warning"

	// A critical condition that requires immediate attention.
	EventSeverity_CRITICAL EventSeverity = "Critical"

)
