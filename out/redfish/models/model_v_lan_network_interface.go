/* -----------------------------------------------------------------
* v_lan_network_interface.go -
*
* DMTF Redfish VLanNetworkInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VLanNetworkInterface - This resource contains information for a Virtual LAN (VLAN) network instance available on a manager, system or other device.
// Modeled after DMTF Redfish schema VLanNetworkInterface
type VLanNetworkInterface struct {
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

	// This indicates if this VLAN is enabled.
	VLANEnable bool `json:"VLANEnable,omitempty"`

	// This indicates the VLAN identifier for this VLAN.
	VLANId VLANId `json:"VLANId,omitempty"`

}
