/* -----------------------------------------------------------------
* enum_flow_control.go -
*
* DMTF Redfish FlowControl enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type FlowControl string

const (
	// No IEEE 802.3x flow control is enabled on this port.
	FlowControl_NONE FlowControl = "None"

	// IEEE 802.3x flow control may be initiated by this station.
	FlowControl_TX FlowControl = "TX"

	// IEEE 802.3x flow control may be initiated by the link partner.
	FlowControl_RX FlowControl = "RX"

	// IEEE 802.3x flow control may be initiated by this station or the link partner.
	FlowControl_TX_RX FlowControl = "TX_RX"

)
