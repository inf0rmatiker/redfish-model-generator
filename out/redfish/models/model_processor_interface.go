/* -----------------------------------------------------------------
* processor_interface.go -
*
* DMTF Redfish ProcessorInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ProcessorInterface - This type describes an interface between the system, or external connection, and the processor.
// Modeled after DMTF Redfish schema ProcessorInterface
type ProcessorInterface struct {
	// The Ethernet-related information for this interface.
	Ethernet EthernetInterface `json:"Ethernet,omitempty"`

	// The interface type.
	InterfaceType SystemInterfaceType `json:"InterfaceType,omitempty"`

	// The PCIe-related information for this interface.
	PCIe map[string]interface{} `json:"PCIe,omitempty"`

}
