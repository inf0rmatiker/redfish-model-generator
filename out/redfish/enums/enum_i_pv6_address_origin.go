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
	// A static address as configured by the user
	IPv6AddressOrigin_STATIC IPv6AddressOrigin = "Static"

	// Address is provided by a DHCPv6 service
	IPv6AddressOrigin_DHC_PV6 IPv6AddressOrigin = "DHCPv6"

	// Address is valid only for this network segment (link)
	IPv6AddressOrigin_LINK_LOCAL IPv6AddressOrigin = "LinkLocal"

	// Address is provided by a Stateless Address AutoConfiguration (SLAAC) service
	IPv6AddressOrigin_SLAAC IPv6AddressOrigin = "SLAAC"

)
