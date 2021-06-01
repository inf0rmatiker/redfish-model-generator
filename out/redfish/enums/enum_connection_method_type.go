/* -----------------------------------------------------------------
* enum_connection_method_type.go -
*
* DMTF Redfish ConnectionMethodType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ConnectionMethodType string

const (
	// Redfish connection method.
	ConnectionMethodType_REDFISH ConnectionMethodType = "Redfish"

	// SNMP connection method.
	ConnectionMethodType_SNMP ConnectionMethodType = "SNMP"

	// IPMI 1.5 connection method.
	ConnectionMethodType_IPMI15 ConnectionMethodType = "IPMI15"

	// IPMI 2.0 connection method.
	ConnectionMethodType_IPMI20 ConnectionMethodType = "IPMI20"

	// NETCONF connection method.
	ConnectionMethodType_NETCONF ConnectionMethodType = "NETCONF"

	// OEM connection method.
	ConnectionMethodType_OEM ConnectionMethodType = "OEM"

)
