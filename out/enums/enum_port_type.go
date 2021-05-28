/* -----------------------------------------------------------------
 * enum_port_type.go -
 *
 * DMTF Redfish PortType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type PortType string

const (
	// This port connects to a host device.
	PortType_UPSTREAM_PORT PortType = "UpstreamPort"

	// This port connects to a target device.
	PortType_DOWNSTREAM_PORT PortType = "DownstreamPort"

	// This port connects to another switch.
	PortType_INTERSWITCH_PORT PortType = "InterswitchPort"

	// This port connects to a switch manager.
	PortType_MANAGEMENT_PORT PortType = "ManagementPort"

	// This port connects to any type of device.
	PortType_BIDIRECTIONAL_PORT PortType = "BidirectionalPort"

	// This port has not yet been configured.
	PortType_UNCONFIGURED_PORT PortType = "UnconfiguredPort"

)
