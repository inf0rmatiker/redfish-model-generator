/* -----------------------------------------------------------------
* enum_account_types.go -
*
* DMTF Redfish AccountTypes enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AccountTypes string

const (
	// Allow access to the Redfish service.
	AccountTypes_REDFISH AccountTypes = "Redfish"

	// Allow access to SNMP services.
	AccountTypes_SNMP AccountTypes = "SNMP"

	// OEM account type.  See the OEMAccountTypes property.
	AccountTypes_OEM AccountTypes = "OEM"

	// Allow access to the host's console, which could be connected through Telnet, SSH, or other protocol.
	AccountTypes_HOST_CONSOLE AccountTypes = "HostConsole"

	// Allow access to the manager's console, which could be connected through Telnet, SSH, SM CLP, or other protocol.
	AccountTypes_MANAGER_CONSOLE AccountTypes = "ManagerConsole"

	// Allow access to the Intelligent Platform Management Interface service.
	AccountTypes_IPMI AccountTypes = "IPMI"

	// Allow access to a Keyboard-Video-Mouse over IP session.
	AccountTypes_KVMIP AccountTypes = "KVMIP"

	// Allow access to control virtual media.
	AccountTypes_VIRTUAL_MEDIA AccountTypes = "VirtualMedia"

	// Allow access to a web user interface session, such as a graphical interface or another web-based protocol.
	AccountTypes_WEB_UI AccountTypes = "WebUI"

)
