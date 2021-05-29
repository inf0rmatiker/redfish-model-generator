/* -----------------------------------------------------------------
* processor_id.go -
*
* DMTF Redfish ProcessorId resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The identification information for a processor.
type ProcessorId struct {
	// The effective family for this processor.
	EffectiveFamily string `json:"EffectiveFamily,omitempty"`

	// The effective model for this processor.
	EffectiveModel string `json:"EffectiveModel,omitempty"`

	// The raw manufacturer-provided processor identification registers for this processor.
	IdentificationRegisters string `json:"IdentificationRegisters,omitempty"`

	// The microcode information for this processor.
	MicrocodeInfo string `json:"MicrocodeInfo,omitempty"`

	// The step value for this processor.
	Step string `json:"Step,omitempty"`

	// The vendor identification for this processor.
	VendorId string `json:"VendorId,omitempty"`

}
