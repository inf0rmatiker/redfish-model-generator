/* -----------------------------------------------------------------
* ssd_protocol.go -
*
* DMTF Redfish SSDProtocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// Indicates if the protocol is enabled or disabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`

	// Indicates the protocol port.
	Port float64 `json:"Port,omitempty"`

	// Indicates how often the Multicast is done from this service for SSDP.
	NotifyMulticastIntervalSeconds float64 `json:"NotifyMulticastIntervalSeconds,omitempty"`

	// Indicates the time to live hop count for SSDPs Notify messages.
	NotifyTTL float64 `json:"NotifyTTL,omitempty"`

	// Indicates the scope for the IPv6 Notify messages for SSDP.
	NotifyIPv6Scope NotifyIPv6Scope `json:"NotifyIPv6Scope,omitempty"`

}
