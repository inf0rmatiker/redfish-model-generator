/* -----------------------------------------------------------------
* pci_id.go -
*
* DMTF Redfish PciId resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PciId - A PCI ID.
// Modeled after DMTF Redfish schema PciId
type PciId struct {
	// The Device ID of this PCIe function.
	DeviceId string `json:"DeviceId,omitempty"`

	// The Subsystem ID of this PCIe function.
	SubsystemId string `json:"SubsystemId,omitempty"`

	// The Subsystem Vendor ID of this PCIe function.
	SubsystemVendorId string `json:"SubsystemVendorId,omitempty"`

	// The Vendor ID of this PCIe function.
	VendorId string `json:"VendorId,omitempty"`

}
