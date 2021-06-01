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
	PinOut_CISCO PinOut = "Cisco"
	PinOut_CYCLADES PinOut = "Cyclades"
	PinOut_DIGI PinOut = "Digi"
)
