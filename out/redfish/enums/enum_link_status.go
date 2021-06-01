/* -----------------------------------------------------------------
* enum_link_status.go -
*
* DMTF Redfish LinkStatus enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type LinkStatus string

const (
	// The port is enabled but link is down.
	LinkStatus_DOWN LinkStatus = "Down"

	// The port is enabled and link is good (up).
	LinkStatus_UP LinkStatus = "Up"

)
