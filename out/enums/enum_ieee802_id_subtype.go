/* -----------------------------------------------------------------
 * enum_ieee802_id_subtype.go -
 *
 * DMTF Redfish IEEE802IdSubtype enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type IEEE802IdSubtype string

const (
	// Chassis component, based in the value of entPhysicalAlias in RFC4133.
	IEEE802IdSubtype_CHASSIS_COMP IEEE802IdSubtype = "ChassisComp"

	// Interface alias, based on the ifAlias MIB object.
	IEEE802IdSubtype_IF_ALIAS IEEE802IdSubtype = "IfAlias"

	// Port component, based in the value of entPhysicalAlias in RFC4133.
	IEEE802IdSubtype_PORT_COMP IEEE802IdSubtype = "PortComp"

	// MAC address, based on an agent detected unicast source address as defined in IEEE Std. 802.
	IEEE802IdSubtype_MAC_ADDR IEEE802IdSubtype = "MacAddr"

	// Network address, based on an agent detected network address.
	IEEE802IdSubtype_NETWORK_ADDR IEEE802IdSubtype = "NetworkAddr"

	// Interface name, based on the ifName MIB object.
	IEEE802IdSubtype_IF_NAME IEEE802IdSubtype = "IfName"

	// Agent circuit ID, based on the agent-local identifier of the circuit as defined in RFC3046.
	IEEE802IdSubtype_AGENT_ID IEEE802IdSubtype = "AgentId"

	// Locally assigned, based on a alpha-numeric value locally assigned.
	IEEE802IdSubtype_LOCAL_ASSIGN IEEE802IdSubtype = "LocalAssign"

	// No data to be sent to/received from remote partner.
	IEEE802IdSubtype_NOT_TRANSMITTED IEEE802IdSubtype = "NotTransmitted"

)
