/* -----------------------------------------------------------------
* v_lan_network_interface.go -
*
* DMTF Redfish VLanNetworkInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The VLanNetworkInterface schema describes a VLAN network instance that is available on a manager, system, or other device.
type VLanNetworkInterface struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether this VLAN is enabled for this interface.
	VLANEnable bool `json:"VLANEnable,omitempty"`

	// The ID for this VLAN.
	VLANId VLANId `json:"VLANId,omitempty"`

	// The priority for this VLAN.
	VLANPriority VLANPriority `json:"VLANPriority,omitempty"`

}
