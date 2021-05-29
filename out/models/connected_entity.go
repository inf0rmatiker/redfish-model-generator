/* -----------------------------------------------------------------
* connected_entity.go -
*
* DMTF Redfish ConnectedEntity resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Represents a remote resource that is connected to the network accessible to this endpoint.
type ConnectedEntity struct {
	// The link to the associated entity.
	EntityLink Resource `json:"EntityLink,omitempty"`

	// The PCI ID of the connected entity.
	EntityPciId PciId `json:"EntityPciId,omitempty"`

	// The role of the connected entity.
	EntityRole EntityRole `json:"EntityRole,omitempty"`

	// The type of the connected entity.
	EntityType EntityType `json:"EntityType,omitempty"`

	// The Gen-Z related properties for the entity.
	GenZ GenZ `json:"GenZ,omitempty"`

	// Identifiers for the remote entity.
	Identifiers array `json:"Identifiers,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The Class Code, Subclass, and Programming Interface code of this PCIe function.
	PciClassCode string `json:"PciClassCode,omitempty"`

	// The PCI ID of the connected entity.
	PciFunctionNumber int `json:"PciFunctionNumber,omitempty"`

}
