/* -----------------------------------------------------------------
* resource_block.go -
*
* DMTF Redfish ResourceBlock resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ResourceBlock - This schema defines a Resource Block resource.
// Modeled after DMTF Redfish schema ResourceBlock
type ResourceBlock struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// This property describes the composition status details for this Resource Block.
	CompositionStatus CompositionStatus `json:"CompositionStatus"`

	// An array of references to the Computer Systems available in this Resource Block.
	ComputerSystems []ComputerSystem `json:"ComputerSystems,omitempty"`

	ComputerSystemsOdataCount int `json:"ComputerSystems@odata.count,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An array of references to the Ethernet Interfaces available in this Resource Block.
	EthernetInterfaces []EthernetInterface `json:"EthernetInterfaces,omitempty"`

	EthernetInterfacesOdataCount int `json:"EthernetInterfaces@odata.count,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An array of references to the Memory available in this Resource Block.
	Memory []Memory `json:"Memory,omitempty"`

	MemoryOdataCount int `json:"Memory@odata.count,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// An array of references to the Network Interfaces available in this Resource Block.
	NetworkInterfaces []NetworkInterface `json:"NetworkInterfaces,omitempty"`

	NetworkInterfacesOdataCount int `json:"NetworkInterfaces@odata.count,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of references to the Processors available in this Resource Block.
	Processors []Processor `json:"Processors,omitempty"`

	ProcessorsOdataCount int `json:"Processors@odata.count,omitempty"`

	// This property represents the types of resources available on this Resource Block.
	ResourceBlockType []ResourceBlockType `json:"ResourceBlockType"`

	// An array of references to the Simple Storage available in this Resource Block.
	SimpleStorage []SimpleStorage `json:"SimpleStorage,omitempty"`

	SimpleStorageOdataCount int `json:"SimpleStorage@odata.count,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// An array of references to the Storage available in this Resource Block.
	Storage []Storage `json:"Storage,omitempty"`

	StorageOdataCount int `json:"Storage@odata.count,omitempty"`

}
