/* -----------------------------------------------------------------
 * enum_graphical_connect_types_supported.go -
 *
 * DMTF Redfish GraphicalConnectTypesSupported enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type GraphicalConnectTypesSupported string

const (
	// The controller supports a graphical console connection through a KVM-IP (redirection of Keyboard, Video, Mouse over IP) protocol.
	GraphicalConnectTypesSupported_KVMIP GraphicalConnectTypesSupported = "KVMIP"

	// The controller supports a graphical console connection through an OEM-specific protocol.
	GraphicalConnectTypesSupported_OEM GraphicalConnectTypesSupported = "Oem"

)
