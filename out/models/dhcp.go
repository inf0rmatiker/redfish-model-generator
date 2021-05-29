/* -----------------------------------------------------------------
* dhcp.go -
*
* DMTF Redfish DHCP resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DHCP related properties for an Ethernet fabric.
type DHCP struct {
	// Dynamic Host Configuration Protocol (DHCP) interface Maximum Transmission Unit (MTU).
	DHCPInterfaceMTUBytes int `json:"DHCPInterfaceMTUBytes,omitempty"`

	// Dynamic Host Configuration Protocol (DHCP) relay status.
	DHCPRelayEnabled bool `json:"DHCPRelayEnabled,omitempty"`

	// The Dynamic Host Configuration Protocol (DHCP) IPv4 addresses for this Ethernet fabric.
	DHCPServer []string `json:"DHCPServer,omitempty"`

}
