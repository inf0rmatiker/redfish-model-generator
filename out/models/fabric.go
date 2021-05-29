/* -----------------------------------------------------------------
* fabric.go -
*
* DMTF Redfish Fabric resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Fabric schema represents a simple fabric consisting of one or more switches, zero or more endpoints, and zero or more zones.
type Fabric struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The collection of links to the address pools that this fabric contains.
	AddressPools AddressPoolCollection `json:"AddressPools,omitempty"`

	// The collection of links to the connections that this fabric contains.
	Connections ConnectionCollection `json:"Connections,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The collection of links to the endpoint groups that this fabric contains.
	EndpointGroups EndpointGroupCollection `json:"EndpointGroups,omitempty"`

	// The collection of links to the endpoints that this fabric contains.
	Endpoints EndpointCollection `json:"Endpoints,omitempty"`

	// The protocol being sent over this fabric.
	FabricType Protocol `json:"FabricType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The maximum number of zones the switch can currently configure.
	MaxZones int `json:"MaxZones,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The collection of links to the switches that this fabric contains.
	Switches SwitchCollection `json:"Switches,omitempty"`

	// The collection of links to the zones that this fabric contains.
	Zones ZoneCollection `json:"Zones,omitempty"`

}
