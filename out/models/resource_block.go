/* -----------------------------------------------------------------
* resource_block.go -
*
* DMTF Redfish ResourceBlock resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ResourceBlock schema contains definitions resource blocks, its components, and affinity to composed devices.
type ResourceBlock struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The client to which this resource block is assigned.
	Client string `json:"Client,omitempty"`

	// The composition status details for this resource block.
	CompositionStatus CompositionStatus `json:"CompositionStatus"`

	// An array of links to the computer systems available in this resource block.
	ComputerSystems array `json:"ComputerSystems,omitempty"`

	ComputerSystems@odata.count count `json:"ComputerSystems@odata.count,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An array of links to the drives available in this resource block.
	Drives array `json:"Drives,omitempty"`

	Drives@odata.count count `json:"Drives@odata.count,omitempty"`

	// An array of links to the Ethernet interfaces available in this resource block.
	EthernetInterfaces array `json:"EthernetInterfaces,omitempty"`

	EthernetInterfaces@odata.count count `json:"EthernetInterfaces@odata.count,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An array of links to the memory available in this resource block.
	Memory array `json:"Memory,omitempty"`

	Memory@odata.count count `json:"Memory@odata.count,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// An array of links to the Network Interfaces available in this resource block.
	NetworkInterfaces array `json:"NetworkInterfaces,omitempty"`

	NetworkInterfaces@odata.count count `json:"NetworkInterfaces@odata.count,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The pool to which this resource block belongs.
	Pool PoolType `json:"Pool,omitempty"`

	// An array of links to the processors available in this resource block.
	Processors array `json:"Processors,omitempty"`

	Processors@odata.count count `json:"Processors@odata.count,omitempty"`

	// The types of resources available on this resource block.
	ResourceBlockType array `json:"ResourceBlockType"`

	// An array of links to the simple storage available in this resource block.
	SimpleStorage array `json:"SimpleStorage,omitempty"`

	SimpleStorage@odata.count count `json:"SimpleStorage@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// An array of links to the storage available in this resource block.
	Storage array `json:"Storage,omitempty"`

	Storage@odata.count count `json:"Storage@odata.count,omitempty"`

}
