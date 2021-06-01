/* -----------------------------------------------------------------
* network_adapter.go -
*
* DMTF Redfish NetworkAdapter resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NetworkAdapter - The NetworkAdapter schema represents a physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
// Modeled after DMTF Redfish schema NetworkAdapter
type NetworkAdapter struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly resource associated with this adapter.
	Assembly map[string]interface{} `json:"Assembly,omitempty"`

	// The set of network controllers ASICs that make up this NetworkAdapter.
	Controllers []Controllers `json:"Controllers,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The manufacturer or OEM of this network adapter.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model string for this network adapter.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The link to the collection of network device functions associated with this network adapter.
	NetworkDeviceFunctions map[string]interface{} `json:"NetworkDeviceFunctions,omitempty"`

	// The link to the collection of network ports associated with this network adapter.
	NetworkPorts map[string]interface{} `json:"NetworkPorts,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Part number for this network adapter.
	PartNumber string `json:"PartNumber,omitempty"`

	// The manufacturer SKU for this network adapter.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this network adapter.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
