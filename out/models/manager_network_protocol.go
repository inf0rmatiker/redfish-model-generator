/* -----------------------------------------------------------------
* manager_network_protocol.go -
*
* DMTF Redfish ManagerNetworkProtocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The network service settings for the manager.
type ManagerNetworkProtocol struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The settings for this manager's DHCPv4 protocol support.
	DHCP Protocol `json:"DHCP,omitempty"`

	// The settings for this manager's DHCPv6 protocol support.
	DHCPv6 Protocol `json:"DHCPv6,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The fully qualified domain name for the manager obtained by DNS including the host name and top-level domain name.
	FQDN string `json:"FQDN,omitempty"`

	// The settings for this manager's HTTP protocol support.
	HTTP Protocol `json:"HTTP,omitempty"`

	// The settings for this manager's HTTPS protocol support.
	HTTPS HTTPSProtocol `json:"HTTPS,omitempty"`

	// The DNS host name of this manager, without any domain information.
	HostName string `json:"HostName,omitempty"`

	// The settings for this manager's IPMI-over-LAN protocol support.
	IPMI Protocol `json:"IPMI,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The settings for this manager's KVM-IP protocol support that apply to all system instances controlled by this manager.
	KVMIP Protocol `json:"KVMIP,omitempty"`

	// The settings for this manager's NTP protocol support.
	NTP NTPProtocol `json:"NTP,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The settings for this manager's Remote Desktop Protocol support.
	RDP Protocol `json:"RDP,omitempty"`

	// The settings for this manager's Remote Frame Buffer protocol support, which can support VNC.
	RFB Protocol `json:"RFB,omitempty"`

	// The settings for this manager's SNMP support.
	SNMP SNMPProtocol `json:"SNMP,omitempty"`

	// The settings for this manager's SSDP support.
	SSDP SSDProtocol `json:"SSDP,omitempty"`

	// The settings for this manager's Secure Shell (SSH) protocol support.
	SSH Protocol `json:"SSH,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

	// The settings for this manager's Telnet protocol support.
	Telnet Protocol `json:"Telnet,omitempty"`

	// The settings for this manager's virtual media support that apply to all system instances controlled by this manager.
	VirtualMedia Protocol `json:"VirtualMedia,omitempty"`

}
