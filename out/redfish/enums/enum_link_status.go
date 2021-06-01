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
	// This link on this interface is up.
	LinkStatus_LINK_UP LinkStatus = "LinkUp"

	// This link on this interface is starting.  A physical link has been established, but the port is not able to transfer data.
	LinkStatus_STARTING LinkStatus = "Starting"

	// This physical link on this interface is training.
	LinkStatus_TRAINING LinkStatus = "Training"

	// The link on this interface is down.
	LinkStatus_LINK_DOWN LinkStatus = "LinkDown"

	// No physical link detected on this interface.
	LinkStatus_NO_LINK LinkStatus = "NoLink"

)
