/* -----------------------------------------------------------------
* fibre_channel.go -
*
* DMTF Redfish FibreChannel resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// FibreChannel - This type describes Fibre Channel capabilities, status, and configuration of a network device function.
// Modeled after DMTF Redfish schema FibreChannel
type FibreChannel struct {
	// Whether the FCoE Initialization Protocol (FIP) is used for populating the FCoE VLAN Id.
	AllowFIPVLANDiscovery bool `json:"AllowFIPVLANDiscovery,omitempty"`

	// An array of Fibre Channel boot targets configured for this network device function.
	BootTargets []BootTargets `json:"BootTargets,omitempty"`

	// The active FCoE VLAN ID.
	FCoEActiveVLANId int `json:"FCoEActiveVLANId,omitempty"`

	// The locally configured FCoE VLAN ID.
	FCoELocalVLANId int `json:"FCoELocalVLANId,omitempty"`

	// This is the permanent WWNN address assigned to this network device function (physical function).
	PermanentWWNN string `json:"PermanentWWNN,omitempty"`

	// This is the permanent WWPN address assigned to this network device function (physical function).
	PermanentWWPN string `json:"PermanentWWPN,omitempty"`

	// This is the currently configured WWNN address of the network device function (physical function).
	WWNN string `json:"WWNN,omitempty"`

	// The configuration source of the WWNs for this connection (WWPN and WWNN).
	WWNSource WWNSource `json:"WWNSource,omitempty"`

	// This is the currently configured WWPN address of the network device function (physical function).
	WWPN string `json:"WWPN,omitempty"`

}
