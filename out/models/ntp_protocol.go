/* -----------------------------------------------------------------
* ntp_protocol.go -
*
* DMTF Redfish NTPProtocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings for a network protocol associated with a manager.
type NTPProtocol struct {
	// Indicates to which NTP servers this manager is subscribed.
	NTPServers []string `json:"NTPServers,omitempty"`

	// The protocol port.
	Port int `json:"Port,omitempty"`

	// An indication of whether the protocol is enabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`

}
