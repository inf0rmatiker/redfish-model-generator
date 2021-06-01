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
	// No flow control imposed
	FlowControl_NONE FlowControl = "None"

	// XON/XOFF in-band flow control imposed
	FlowControl_SOFTWARE FlowControl = "Software"

	// Out of band flow control imposed
	FlowControl_HARDWARE FlowControl = "Hardware"

)
