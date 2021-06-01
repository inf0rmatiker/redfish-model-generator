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
	// A static address as configured by the user
	IPv4AddressOrigin_STATIC IPv4AddressOrigin = "Static"

	// Address is provided by a DHCPv4 service
	IPv4AddressOrigin_DHCP IPv4AddressOrigin = "DHCP"

	// Address is provided by a BOOTP service
	IPv4AddressOrigin_BOOTP IPv4AddressOrigin = "BOOTP"

	// Address is valid only for this network segment (link)
	IPv4AddressOrigin_I_PV4_LINK_LOCAL IPv4AddressOrigin = "IPv4LinkLocal"

)
