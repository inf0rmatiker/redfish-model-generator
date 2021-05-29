/* -----------------------------------------------------------------
* network_interface.go -
*
* DMTF Redfish NetworkInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The NetworkInterface schema describes links to the network adapters, network ports, and network device functions, and represents the functionality available to the containing system.
type NetworkInterface struct {
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

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The link to the network device functions associated with this network interface.
	NetworkDeviceFunctions NetworkDeviceFunctionCollection `json:"NetworkDeviceFunctions,omitempty"`

	// The link to the network ports associated with this network interface.
	NetworkPorts NetworkPortCollection `json:"NetworkPorts,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to the ports associated with this network interface.
	Ports PortCollection `json:"Ports,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
