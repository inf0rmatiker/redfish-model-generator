/* -----------------------------------------------------------------
* pci_id.go -
*
* DMTF Redfish PciId resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A PCI ID.
type PciId struct {
	// The Class Code, Subclass, and Programming Interface code of this PCIe function.
	ClassCode string `json:"ClassCode,omitempty"`

	// The Device ID of this PCIe function.
	DeviceId string `json:"DeviceId,omitempty"`

	// The PCI ID of the connected entity.
	FunctionNumber int `json:"FunctionNumber,omitempty"`

	// The Subsystem ID of this PCIe function.
	SubsystemId string `json:"SubsystemId,omitempty"`

	// The Subsystem Vendor ID of this PCIe function.
	SubsystemVendorId string `json:"SubsystemVendorId,omitempty"`

	// The Vendor ID of this PCIe function.
	VendorId string `json:"VendorId,omitempty"`

}
