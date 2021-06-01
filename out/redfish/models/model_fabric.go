/* -----------------------------------------------------------------
* fabric.go -
*
* DMTF Redfish Fabric resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Fabric - Fabric contains properties describing a simple fabric consisting of one or more switches, zero or more endpoints, and zero or more zones.
// Modeled after DMTF Redfish schema Fabric
type Fabric struct {
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

	// The protocol being sent over this fabric.
	FabricType map[string]interface{} `json:"FabricType,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

	// The value of this property shall contain the maximum number of zones the switch can currently configure.
	MaxZones float64 `json:"MaxZones,omitempty"`

	// Contains references to other resources that are related to this resource.
	Links map[string]interface{} `json:"Links,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// A collection of references to the zones contained in this fabric.
	Zones map[string]interface{} `json:"Zones,omitempty"`

	// A collection of references to the endpoints contained in this fabric.
	Endpoints map[string]interface{} `json:"Endpoints,omitempty"`

	// A collection of references to the switches contained in this fabric.
	Switches map[string]interface{} `json:"Switches,omitempty"`

}
