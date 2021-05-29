/* -----------------------------------------------------------------
* i_scsi_boot.go -
*
* DMTF Redfish iSCSIBoot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The iSCSI boot capabilities, status, and configuration for a network device function.
type iSCSIBoot struct {
	// The iSCSI boot authentication method for this network device function.
	AuthenticationMethod AuthenticationMethod `json:"AuthenticationMethod,omitempty"`

	// The shared secret for CHAP authentication.
	CHAPSecret string `json:"CHAPSecret,omitempty"`

	// The user name for CHAP authentication.
	CHAPUsername string `json:"CHAPUsername,omitempty"`

	// The type of IP address being populated in the iSCSIBoot IP address fields.
	IPAddressType IPAddressType `json:"IPAddressType,omitempty"`

	// An indication of whether the iSCSI boot initiator uses DHCP to obtain the initiator name, IP address, and netmask.
	IPMaskDNSViaDHCP bool `json:"IPMaskDNSViaDHCP,omitempty"`

	// The IPv6 or IPv4 iSCSI boot default gateway.
	InitiatorDefaultGateway string `json:"InitiatorDefaultGateway,omitempty"`

	// The IPv6 or IPv4 address of the iSCSI initiator.
	InitiatorIPAddress string `json:"InitiatorIPAddress,omitempty"`

	// The iSCSI initiator name.
	InitiatorName string `json:"InitiatorName,omitempty"`

	// The IPv6 or IPv4 netmask of the iSCSI boot initiator.
	InitiatorNetmask string `json:"InitiatorNetmask,omitempty"`

	// The CHAP secret for two-way CHAP authentication.
	MutualCHAPSecret string `json:"MutualCHAPSecret,omitempty"`

	// The CHAP user name for two-way CHAP authentication.
	MutualCHAPUsername string `json:"MutualCHAPUsername,omitempty"`

	// The IPv6 or IPv4 address of the primary DNS server for the iSCSI boot initiator.
	PrimaryDNS string `json:"PrimaryDNS,omitempty"`

	// The logical unit number (LUN) for the primary iSCSI boot target.
	PrimaryLUN int `json:"PrimaryLUN,omitempty"`

	// The IPv4 or IPv6 address for the primary iSCSI boot target.
	PrimaryTargetIPAddress string `json:"PrimaryTargetIPAddress,omitempty"`

	// The name of the iSCSI primary boot target.
	PrimaryTargetName string `json:"PrimaryTargetName,omitempty"`

	// The TCP port for the primary iSCSI boot target.
	PrimaryTargetTCPPort int `json:"PrimaryTargetTCPPort,omitempty"`

	// An indication of whether the primary VLAN is enabled.
	PrimaryVLANEnable bool `json:"PrimaryVLANEnable,omitempty"`

	// The 802.1q VLAN ID to use for iSCSI boot from the primary target.
	PrimaryVLANId int `json:"PrimaryVLANId,omitempty"`

	// An indication of whether IPv6 router advertisement is enabled for the iSCSI boot target.
	RouterAdvertisementEnabled bool `json:"RouterAdvertisementEnabled,omitempty"`

	// The IPv6 or IPv4 address of the secondary DNS server for the iSCSI boot initiator.
	SecondaryDNS string `json:"SecondaryDNS,omitempty"`

	// The logical unit number (LUN) for the secondary iSCSI boot target.
	SecondaryLUN int `json:"SecondaryLUN,omitempty"`

	// The IPv4 or IPv6 address for the secondary iSCSI boot target.
	SecondaryTargetIPAddress string `json:"SecondaryTargetIPAddress,omitempty"`

	// The name of the iSCSI secondary boot target.
	SecondaryTargetName string `json:"SecondaryTargetName,omitempty"`

	// The TCP port for the secondary iSCSI boot target.
	SecondaryTargetTCPPort int `json:"SecondaryTargetTCPPort,omitempty"`

	// An indication of whether the secondary VLAN is enabled.
	SecondaryVLANEnable bool `json:"SecondaryVLANEnable,omitempty"`

	// The 802.1q VLAN ID to use for iSCSI boot from the secondary target.
	SecondaryVLANId int `json:"SecondaryVLANId,omitempty"`

	// An indication of whether the iSCSI boot target name, LUN, IP address, and netmask should be obtained from DHCP.
	TargetInfoViaDHCP bool `json:"TargetInfoViaDHCP,omitempty"`

}
