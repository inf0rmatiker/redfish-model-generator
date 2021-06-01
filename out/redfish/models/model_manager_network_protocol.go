/* -----------------------------------------------------------------
* manager_network_protocol.go -
*
* DMTF Redfish ManagerNetworkProtocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ManagerNetworkProtocol - This resource is used to obtain or modify the network services managed by a given manager.
// Modeled after DMTF Redfish schema ManagerNetworkProtocol
type ManagerNetworkProtocol struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The DNS Host Name of this manager, without any domain information.
	HostName string `json:"HostName,omitempty"`

	// This is the fully qualified domain name for the manager obtained by DNS including the host name and top-level domain name.
	FQDN string `json:"FQDN,omitempty"`

	// Settings for this Manager's HTTP protocol support.
	HTTP Protocol `json:"HTTP,omitempty"`

	// Settings for this Manager's HTTPS protocol support.
	HTTPS Protocol `json:"HTTPS,omitempty"`

	// Settings for this Manager's SNMP support.
	SNMP Protocol `json:"SNMP,omitempty"`

	// Settings for this Manager's Virtual Media support.
	VirtualMedia Protocol `json:"VirtualMedia,omitempty"`

	// Settings for this Manager's Telnet protocol support.
	Telnet Protocol `json:"Telnet,omitempty"`

	// Settings for this Manager's SSDP support.
	SSDP SSDProtocol `json:"SSDP,omitempty"`

	// Settings for this Manager's IPMI-over-LAN protocol support.
	IPMI Protocol `json:"IPMI,omitempty"`

	// Settings for this Manager's SSH (Secure Shell) protocol support.
	SSH Protocol `json:"SSH,omitempty"`

	// Settings for this Manager's KVM-IP protocol support.
	KVMIP Protocol `json:"KVMIP,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

}
