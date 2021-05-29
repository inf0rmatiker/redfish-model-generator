/* -----------------------------------------------------------------
* sfp.go -
*
* DMTF Redfish SFP resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A small form-factor pluggable (SFP) device attached to a port.
type SFP struct {
	// The type of fiber connection currently used by this SFP.
	FiberConnectionType FiberConnectionType `json:"FiberConnectionType,omitempty"`

	// The manufacturer of this SFP.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The medium type connected to this SFP.
	MediumType MediumType `json:"MediumType,omitempty"`

	// The part number for this SFP.
	PartNumber string `json:"PartNumber,omitempty"`

	// The serial number for this SFP.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The types of SFP devices that can be attached to this port.
	SupportedSFPTypes array `json:"SupportedSFPTypes,omitempty"`

	// The type of SFP device that is attached to this port.
	Type SFPType `json:"Type,omitempty"`

}
