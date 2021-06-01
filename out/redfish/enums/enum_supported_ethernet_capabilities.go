/* -----------------------------------------------------------------
* enum_supported_ethernet_capabilities.go -
*
* DMTF Redfish SupportedEthernetCapabilities enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SupportedEthernetCapabilities string

const (
	// Wake on LAN (WoL) is supported on this port.
	SupportedEthernetCapabilities_WAKE_ON_LAN SupportedEthernetCapabilities = "WakeOnLAN"

	// IEEE 802.3az Energy Efficient Ethernet (EEE) is supported on this port.
	SupportedEthernetCapabilities_EEE SupportedEthernetCapabilities = "EEE"

)
