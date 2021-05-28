/* -----------------------------------------------------------------
 * enum_ip_address_type.go -
 *
 * DMTF Redfish IPAddressType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type IPAddressType string

const (
	// IPv4 addressing is used for all IP-fields in this object.
	IPAddressType_I_PV4 IPAddressType = "IPv4"

	// IPv6 addressing is used for all IP-fields in this object.
	IPAddressType_I_PV6 IPAddressType = "IPv6"

)
