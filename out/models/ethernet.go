/* -----------------------------------------------------------------
* ethernet.go -
*
* DMTF Redfish Ethernet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Ethernet related properties for an address pool.
type Ethernet struct {
	// Bidirectional Forwarding Detection (BFD) related properties for this Ethernet fabric.
	BFDSingleHopOnly BFDSingleHopOnly `json:"BFDSingleHopOnly,omitempty"`

	// BGP Ethernet Virtual Private Network (EVPN) related properties for this Ethernet fabric.
	BGPEvpn BGPEvpn `json:"BGPEvpn,omitempty"`

	// External BGP (eBGP) related properties for this Ethernet fabric.
	EBGP EBGP `json:"EBGP,omitempty"`

	// IPv4 and Virtual LAN (VLAN) related addressing for this Ethernet fabric.
	IPv4 IPv4 `json:"IPv4,omitempty"`

	// Multi Protocol eBGP (MP eBGP) related properties for this Ethernet fabric.
	MultiProtocolEBGP EBGP `json:"MultiProtocolEBGP,omitempty"`

	// Multi Protocol iBGP (MP iBGP) related properties for this Ethernet fabric.
	MultiProtocolIBGP CommonBGPProperties `json:"MultiProtocolIBGP,omitempty"`

}
