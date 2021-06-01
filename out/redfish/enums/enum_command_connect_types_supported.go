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
	// The controller supports a Command Shell connection using the SSH protocol.
	CommandConnectTypesSupported_SSH CommandConnectTypesSupported = "SSH"

	// The controller supports a Command Shell connection using the Telnet protocol.
	CommandConnectTypesSupported_TELNET CommandConnectTypesSupported = "Telnet"

	// The controller supports a Command Shell connection using the IPMI Serial-over-LAN (SOL) protocol.
	CommandConnectTypesSupported_IPMI CommandConnectTypesSupported = "IPMI"

	// The controller supports a Command Shell connection using an OEM-specific protocol.
	CommandConnectTypesSupported_OEM CommandConnectTypesSupported = "Oem"

)
