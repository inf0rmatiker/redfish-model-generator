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
	SignalType_RS232 SignalType = "Rs232"
	SignalType_RS485 SignalType = "Rs485"
)
