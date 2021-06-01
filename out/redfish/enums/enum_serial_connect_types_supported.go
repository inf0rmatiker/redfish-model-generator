/* -----------------------------------------------------------------
* enum_serial_connect_types_supported.go -
*
* DMTF Redfish SerialConnectTypesSupported enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SerialConnectTypesSupported string

const (
	// The controller supports a Serial Console connection using the SSH protocol.
	SerialConnectTypesSupported_SSH SerialConnectTypesSupported = "SSH"

	// The controller supports a Serial Console connection using the Telnet protocol.
	SerialConnectTypesSupported_TELNET SerialConnectTypesSupported = "Telnet"

	// The controller supports a Serial Console connection using the IPMI Serial-over-LAN (SOL) protocol.
	SerialConnectTypesSupported_IPMI SerialConnectTypesSupported = "IPMI"

	// The controller supports a Serial Console connection using an OEM-specific protocol.
	SerialConnectTypesSupported_OEM SerialConnectTypesSupported = "Oem"

)