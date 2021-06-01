/* -----------------------------------------------------------------
* connected_entity.go -
*
* DMTF Redfish ConnectedEntity resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ConnectedEntity - Represents a remote resource that is connected to the network accessible to this endpoint.
// Modeled after DMTF Redfish schema ConnectedEntity
type ConnectedEntity struct {
	// A link to the associated entity.
	EntityLink map[string]interface{} `json:"EntityLink,omitempty"`

	// The PCI ID of the connected entity.
	EntityPciId PciId `json:"EntityPciId,omitempty"`

	// The role of the connected entity.
	EntityRole EntityRole `json:"EntityRole,omitempty"`

	// The type of the connected entity.
	EntityType EntityType `json:"EntityType,omitempty"`

	// Identifiers for the remote entity.
	Identifiers []Identifier `json:"Identifiers,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The Class Code, Subclass code, and Programming Interface code of this PCIe function.
	PciClassCode string `json:"PciClassCode,omitempty"`

	// The PCI ID of the connected entity.
	PciFunctionNumber float64 `json:"PciFunctionNumber,omitempty"`

}
