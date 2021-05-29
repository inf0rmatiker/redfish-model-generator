/* -----------------------------------------------------------------
* endpoint.go -
*
* DMTF Redfish Endpoint resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Endpoint schema contains the properties of an endpoint resource that represents the properties of an entity that sends or receives protocol-defined messages over a transport.
type Endpoint struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// All the entities connected to this endpoint.
	ConnectedEntities array `json:"ConnectedEntities,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The protocol supported by this endpoint.
	EndpointProtocol Protocol `json:"EndpointProtocol,omitempty"`

	// The amount of memory in bytes that the host should allocate to connect to this endpoint.
	HostReservationMemoryBytes int `json:"HostReservationMemoryBytes,omitempty"`

	// An array of details for each IP transport supported by this endpoint.  The array structure can model multiple IP addresses for this endpoint.
	IPTransportDetails array `json:"IPTransportDetails,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Identifiers for this endpoint.
	Identifiers array `json:"Identifiers,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The PCI ID of the endpoint.
	PciId PciId `json:"PciId,omitempty"`

	// Redundancy information for the lower-level endpoints supporting this endpoint.
	Redundancy array `json:"Redundancy,omitempty"`

	Redundancy@odata.count count `json:"Redundancy@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
