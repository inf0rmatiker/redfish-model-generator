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
	// The link is available for communication on this interface.
	LinkStatus_LINK_UP LinkStatus = "LinkUp"

	// No link or connection is detected on this interface.
	LinkStatus_NO_LINK LinkStatus = "NoLink"

	// No link is detected on this interface, but the interface is connected.
	LinkStatus_LINK_DOWN LinkStatus = "LinkDown"

)
