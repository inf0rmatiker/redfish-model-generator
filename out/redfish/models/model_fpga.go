/* -----------------------------------------------------------------
* fpga.go -
*
* DMTF Redfish FPGA resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// FPGA - The properties of the FPGA device.
// Modeled after DMTF Redfish schema FPGA
type FPGA struct {
	// An array of the FPGA external interfaces.
	ExternalInterfaces []ProcessorInterface `json:"ExternalInterfaces,omitempty"`

	// The FPGA firmware identifier.
	FirmwareId string `json:"FirmwareId,omitempty"`

	// The FPGA firmware manufacturer.
	FirmwareManufacturer string `json:"FirmwareManufacturer,omitempty"`

	// The FPGA firmware version.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The FPGA type.
	FpgaType FpgaType `json:"FpgaType,omitempty"`

	// The FPGA interface to the host.
	HostInterface ProcessorInterface `json:"HostInterface,omitempty"`

	// The FPGA model.
	Model string `json:"Model,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The number of the PCIe Virtual Functions.
	PCIeVirtualFunctions int `json:"PCIeVirtualFunctions,omitempty"`

	// An indication of whether the FPGA firmware can be reprogrammed from the host by using system software.
	ProgrammableFromHost bool `json:"ProgrammableFromHost,omitempty"`

	// An array of the FPGA reconfiguration slots.  An FPGA uses a reconfiguration slot to contain an acceleration function that can change as the FPGA is provisioned.
	ReconfigurationSlots []FpgaReconfigurationSlot `json:"ReconfigurationSlots,omitempty"`

}
