/* -----------------------------------------------------------------
 * enum_i_pv4_address_origin.go -
 *
 * DMTF Redfish IPv4AddressOrigin enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type IPv4AddressOrigin string

const (
	// A user-configured static address.
	IPv4AddressOrigin_STATIC IPv4AddressOrigin = "Static"

	// A DHCPv4 service-provided address.
	IPv4AddressOrigin_DHCP IPv4AddressOrigin = "DHCP"

	// A BOOTP service-provided address.
	IPv4AddressOrigin_BOOTP IPv4AddressOrigin = "BOOTP"

	// The address is valid for only this network segment, or link.
	IPv4AddressOrigin_I_PV4_LINK_LOCAL IPv4AddressOrigin = "IPv4LinkLocal"

)
