/* -----------------------------------------------------------------
 * enum_stop_bits.go -
 *
 * DMTF Redfish StopBits enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type StopBits string

const (
	// One stop bit following the data bits.
	StopBits_1 StopBits = "1"

	// Two stop bits following the data bits.
	StopBits_2 StopBits = "2"

)
