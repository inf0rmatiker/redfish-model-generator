/* -----------------------------------------------------------------
* enum_connected_via.go -
*
* DMTF Redfish ConnectedVia enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ConnectedVia string

const (
	// No current connection
	ConnectedVia_NOT_CONNECTED ConnectedVia = "NotConnected"

	// Connected to a URI location
	ConnectedVia_URI ConnectedVia = "URI"

	// Connected to a client application
	ConnectedVia_APPLET ConnectedVia = "Applet"

	// Connected via an OEM-defined method
	ConnectedVia_OEM ConnectedVia = "Oem"

)
