/* -----------------------------------------------------------------
* endpoint.go -
*
* DMTF Redfish Endpoint resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Endpoint - This is the schema definition for the Endpoint resource. It represents the properties of an entity that sends or receives protocol defined messages over a transport.
// Modeled after DMTF Redfish schema Endpoint
type Endpoint struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The Actions object contains the available custom actions on this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// All the entities connected to this endpoint.
	ConnectedEntities []ConnectedEntity `json:"ConnectedEntities,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The protocol supported by this endpoint.
	EndpointProtocol map[string]interface{} `json:"EndpointProtocol,omitempty"`

	// The amount of memory in Bytes that the Host should allocate to connect to this endpoint.
	HostReservationMemoryBytes float64 `json:"HostReservationMemoryBytes,omitempty"`

	// This array contains details for each IP transport supported by this endpoint. The array structure can be used to model multiple IP addresses for this endpoint.
	IPTransportDetails []IPTransportDetails `json:"IPTransportDetails,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Identifiers for this endpoint.
	Identifiers []Identifier `json:"Identifiers,omitempty"`

	// The links object contains the links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The PCI ID of the endpoint.
	PciId PciId `json:"PciId,omitempty"`

	// Redundancy information for the lower level endpoints supporting this endpoint.
	Redundancy []Redundancy `json:"Redundancy,omitempty"`

	RedundancyOdataCount int `json:"Redundancy@odata.count,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

}
