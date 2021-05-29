/* -----------------------------------------------------------------
* bgp_neighbor.go -
*
* DMTF Redfish BGPNeighbor resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Border Gateway Protocol (BGP) neighbor related properties.
type BGPNeighbor struct {
	// Border Gateway Protocol (BGP) neighbor address.
	Address string `json:"Address,omitempty"`

	// Allow own Autonomous System (AS) status.
	AllowOwnASEnabled bool `json:"AllowOwnASEnabled,omitempty"`

	// Border Gateway Protocol (BGP) retry timer in seconds.
	ConnectRetrySeconds int `json:"ConnectRetrySeconds,omitempty"`

	// Border Gateway Protocol (BGP) hold timer in seconds.
	HoldTimeSeconds int `json:"HoldTimeSeconds,omitempty"`

	// Border Gateway Protocol (BGP) Keepalive timer in seconds.
	KeepaliveIntervalSeconds int `json:"KeepaliveIntervalSeconds,omitempty"`

	// Local Autonomous System (AS) number.
	LocalAS int `json:"LocalAS,omitempty"`

	// Border Gateway Protocol (BGP) neighbor log state change status.
	LogStateChangesEnabled bool `json:"LogStateChangesEnabled,omitempty"`

	// Border Gateway Protocol (BGP) max prefix properties.
	MaxPrefix MaxPrefix `json:"MaxPrefix,omitempty"`

	// Minimum Border Gateway Protocol (BGP) advertisement interval in seconds.
	MinimumAdvertisementIntervalSeconds int `json:"MinimumAdvertisementIntervalSeconds,omitempty"`

	// Border Gateway Protocol (BGP) passive mode status.
	PassiveModeEnabled bool `json:"PassiveModeEnabled,omitempty"`

	// Path MTU discovery status.
	PathMTUDiscoveryEnabled bool `json:"PathMTUDiscoveryEnabled,omitempty"`

	// Peer Autonomous System (AS) number.
	PeerAS int `json:"PeerAS,omitempty"`

	// Replace Border Gateway Protocol (BGP) peer Autonomous System (AS) status.
	ReplacePeerASEnabled bool `json:"ReplacePeerASEnabled,omitempty"`

	// TCP max segment size in Bytes.
	TCPMaxSegmentSizeBytes int `json:"TCPMaxSegmentSizeBytes,omitempty"`

	// Border Gateway Protocol (BGP) treat as withdraw status.
	TreatAsWithdrawEnabled bool `json:"TreatAsWithdrawEnabled,omitempty"`

}
