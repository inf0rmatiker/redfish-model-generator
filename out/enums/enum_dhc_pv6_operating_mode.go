/* -----------------------------------------------------------------
 * enum_dhc_pv6_operating_mode.go -
 *
 * DMTF Redfish DHCPv6OperatingMode enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DHCPv6OperatingMode string

const (
	// DHCPv6 stateful mode.
	DHCPv6OperatingMode_STATEFUL DHCPv6OperatingMode = "Stateful"

	// DHCPv6 stateless mode.
	DHCPv6OperatingMode_STATELESS DHCPv6OperatingMode = "Stateless"

	// DHCPv6 is disabled.
	DHCPv6OperatingMode_DISABLED DHCPv6OperatingMode = "Disabled"

)
