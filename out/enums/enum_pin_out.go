/* -----------------------------------------------------------------
 * enum_pin_out.go -
 *
 * DMTF Redfish PinOut enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type PinOut string

const (
	// The Cisco pinout configuration.
	PinOut_CISCO PinOut = "Cisco"

	// The Cyclades pinout configuration.
	PinOut_CYCLADES PinOut = "Cyclades"

	// The Digi pinout configuration.
	PinOut_DIGI PinOut = "Digi"

)
