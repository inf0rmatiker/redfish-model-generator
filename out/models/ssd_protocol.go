/* -----------------------------------------------------------------
* ssd_protocol.go -
*
* DMTF Redfish SSDProtocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings for a network protocol associated with a manager.
type SSDProtocol struct {
	// The IPv6 scope for multicast NOTIFY messages for SSDP.
	NotifyIPv6Scope NotifyIPv6Scope `json:"NotifyIPv6Scope,omitempty"`

	// The time interval, in seconds, between transmissions of the multicast NOTIFY ALIVE message from this service for SSDP.
	NotifyMulticastIntervalSeconds int `json:"NotifyMulticastIntervalSeconds,omitempty"`

	// The time-to-live hop count for SSDP multicast NOTIFY messages.
	NotifyTTL int `json:"NotifyTTL,omitempty"`

	// The protocol port.
	Port int `json:"Port,omitempty"`

	// An indication of whether the protocol is enabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`

}
