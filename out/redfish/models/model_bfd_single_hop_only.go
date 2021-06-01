/* -----------------------------------------------------------------
* bfd_single_hop_only.go -
*
* DMTF Redfish BFDSingleHopOnly resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// BFDSingleHopOnly - Bidirectional Forwarding Detection (BFD) related properties for an Ethernet fabric.
// Modeled after DMTF Redfish schema BFDSingleHopOnly
type BFDSingleHopOnly struct {
	// Bidirectional Forwarding Detection (BFD) Demand Mode status.
	DemandModeEnabled bool `json:"DemandModeEnabled,omitempty"`

	// Desired Bidirectional Forwarding Detection (BFD) minimal transmit interval.
	DesiredMinTxIntervalMilliseconds int `json:"DesiredMinTxIntervalMilliseconds,omitempty"`

	// Bidirectional Forwarding Detection (BFD) Key Chain name.
	KeyChain string `json:"KeyChain,omitempty"`

	// Bidirectional Forwarding Detection (BFD) multiplier value.
	LocalMultiplier int `json:"LocalMultiplier,omitempty"`

	// Meticulous MD5 authentication of the Bidirectional Forwarding Detection (BFD) session.
	MeticulousModeEnabled bool `json:"MeticulousModeEnabled,omitempty"`

	// Bidirectional Forwarding Detection (BFD) receive value.
	RequiredMinRxIntervalMilliseconds int `json:"RequiredMinRxIntervalMilliseconds,omitempty"`

	// Bidirectional Forwarding Detection (BFD) source port.
	SourcePort int `json:"SourcePort,omitempty"`

}
