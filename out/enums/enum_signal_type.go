/* -----------------------------------------------------------------
 * enum_signal_type.go -
 *
 * DMTF Redfish SignalType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SignalType string

const (
	// The serial interface follows RS232.
	SignalType_RS232 SignalType = "Rs232"

	// The serial interface follows RS485.
	SignalType_RS485 SignalType = "Rs485"

)
