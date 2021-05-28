/* -----------------------------------------------------------------
 * enum_session_types.go -
 *
 * DMTF Redfish SessionTypes enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SessionTypes string

const (
	// The host's console, which could be connected through Telnet, SSH, or other protocol.
	SessionTypes_HOST_CONSOLE SessionTypes = "HostConsole"

	// The manager's console, which could be connected through Telnet, SSH, SM CLP, or other protocol.
	SessionTypes_MANAGER_CONSOLE SessionTypes = "ManagerConsole"

	// Intelligent Platform Management Interface.
	SessionTypes_IPMI SessionTypes = "IPMI"

	// Keyboard-Video-Mouse over IP Session.
	SessionTypes_KVMIP SessionTypes = "KVMIP"

	// OEM Type.  For OEM session types, see the OemSessionType property.
	SessionTypes_OEM SessionTypes = "OEM"

	// A Redfish session.
	SessionTypes_REDFISH SessionTypes = "Redfish"

	// Virtual media.
	SessionTypes_VIRTUAL_MEDIA SessionTypes = "VirtualMedia"

	// A non-Redfish web user interface session, such as a graphical interface or another web-based protocol.
	SessionTypes_WEB_UI SessionTypes = "WebUI"

)
