/* -----------------------------------------------------------------
* fibre_channel.go -
*
* DMTF Redfish FibreChannel resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes Fibre Channel capabilities, status, and configuration for a network device function.
type FibreChannel struct {
	// An indication of whether the FCoE Initialization Protocol (FIP) populates the FCoE VLAN ID.
	AllowFIPVLANDiscovery bool `json:"AllowFIPVLANDiscovery,omitempty"`

	// An array of Fibre Channel boot targets configured for this network device function.
	BootTargets array `json:"BootTargets,omitempty"`

	// The active FCoE VLAN ID.
	FCoEActiveVLANId int `json:"FCoEActiveVLANId,omitempty"`

	// The locally configured FCoE VLAN ID.
	FCoELocalVLANId int `json:"FCoELocalVLANId,omitempty"`

	// The Fibre Channel ID that the switch assigns for this interface.
	FibreChannelId string `json:"FibreChannelId,omitempty"`

	// The permanent World Wide Node Name (WWNN) address assigned to this function.
	PermanentWWNN string `json:"PermanentWWNN,omitempty"`

	// The permanent World Wide Port Name (WWPN) address assigned to this function.
	PermanentWWPN string `json:"PermanentWWPN,omitempty"`

	// The currently configured World Wide Node Name (WWNN) address of this function.
	WWNN string `json:"WWNN,omitempty"`

	// The configuration source of the World Wide Names (WWN) for this World Wide Node Name (WWNN) and World Wide Port Name (WWPN) connection.
	WWNSource WWNSource `json:"WWNSource,omitempty"`

	// The currently configured World Wide Port Name (WWPN) address of this function.
	WWPN string `json:"WWPN,omitempty"`

}
