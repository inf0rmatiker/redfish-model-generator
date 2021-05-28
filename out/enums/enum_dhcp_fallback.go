/* -----------------------------------------------------------------
 * enum_dhcp_fallback.go -
 *
 * DMTF Redfish DHCPFallback enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DHCPFallback string

const (
	// Fall back to a static address specified by IPv4StaticAddresses.
	DHCPFallback_STATIC DHCPFallback = "Static"

	// Fall back to an autoconfigured address.
	DHCPFallback_AUTO_CONFIG DHCPFallback = "AutoConfig"

	// Continue attempting DHCP without a fallback address.
	DHCPFallback_NONE DHCPFallback = "None"

)
