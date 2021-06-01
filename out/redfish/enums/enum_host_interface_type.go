/* -----------------------------------------------------------------
* enum_host_interface_type.go -
*
* DMTF Redfish HostInterfaceType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type HostInterfaceType string

const (
	// This interface is a Network Host Interface.
	HostInterfaceType_NETWORK_HOST_INTERFACE HostInterfaceType = "NetworkHostInterface"

)
