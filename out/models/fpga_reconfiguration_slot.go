/* -----------------------------------------------------------------
* fpga_reconfiguration_slot.go -
*
* DMTF Redfish FpgaReconfigurationSlot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes the FPGA reconfiguration slot.  An FPGA uses a reconfiguration slot to contain an acceleration function that can change as the FPGA is provisioned.
type FpgaReconfigurationSlot struct {
	// The link to the acceleration function that the code programmed into a reconfiguration slot provides.
	AccelerationFunction AccelerationFunction `json:"AccelerationFunction,omitempty"`

	// An indication of whether the reconfiguration slot can be reprogrammed from the host by using system software.
	ProgrammableFromHost bool `json:"ProgrammableFromHost,omitempty"`

	// The FPGA reconfiguration slot identifier.
	SlotId string `json:"SlotId,omitempty"`

	// The UUID for this reconfiguration slot.
	UUID UUID `json:"UUID,omitempty"`

}
