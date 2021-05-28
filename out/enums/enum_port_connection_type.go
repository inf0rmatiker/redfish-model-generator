/* -----------------------------------------------------------------
 * enum_port_connection_type.go -
 *
 * DMTF Redfish PortConnectionType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type PortConnectionType string

const (
	// This port is not connected.
	PortConnectionType_NOT_CONNECTED PortConnectionType = "NotConnected"

	// This port connects through an N-Port to a switch.
	PortConnectionType_N_PORT PortConnectionType = "NPort"

	// This port connects in a Point-to-point configuration.
	PortConnectionType_POINT_TO_POINT PortConnectionType = "PointToPoint"

	// This port connects in a private loop configuration.
	PortConnectionType_PRIVATE_LOOP PortConnectionType = "PrivateLoop"

	// This port connects in a public configuration.
	PortConnectionType_PUBLIC_LOOP PortConnectionType = "PublicLoop"

	// This port connection type is a generic fabric port.
	PortConnectionType_GENERIC PortConnectionType = "Generic"

	// This port connection type is an extender fabric port.
	PortConnectionType_EXTENDER_FABRIC PortConnectionType = "ExtenderFabric"

)
