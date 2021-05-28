/* -----------------------------------------------------------------
 * enum_i_pv6_address_origin.go -
 *
 * DMTF Redfish IPv6AddressOrigin enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type IPv6AddressOrigin string

const (
	// A static user-configured address.
	IPv6AddressOrigin_STATIC IPv6AddressOrigin = "Static"

	// A DHCPv6 service-provided address.
	IPv6AddressOrigin_DHC_PV6 IPv6AddressOrigin = "DHCPv6"

	// The address is valid for only this network segment, or link.
	IPv6AddressOrigin_LINK_LOCAL IPv6AddressOrigin = "LinkLocal"

	// A stateless autoconfiguration (SLAAC) service-provided address.
	IPv6AddressOrigin_SLAAC IPv6AddressOrigin = "SLAAC"

)
