/* -----------------------------------------------------------------
 * enum_command_connect_types_supported.go -
 *
 * DMTF Redfish CommandConnectTypesSupported enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type CommandConnectTypesSupported string

const (
	// The controller supports a command shell connection through the SSH protocol.
	CommandConnectTypesSupported_SSH CommandConnectTypesSupported = "SSH"

	// The controller supports a command shell connection through the Telnet protocol.
	CommandConnectTypesSupported_TELNET CommandConnectTypesSupported = "Telnet"

	// The controller supports a command shell connection through the IPMI Serial Over LAN (SOL) protocol.
	CommandConnectTypesSupported_IPMI CommandConnectTypesSupported = "IPMI"

	// The controller supports a command shell connection through an OEM-specific protocol.
	CommandConnectTypesSupported_OEM CommandConnectTypesSupported = "Oem"

)
