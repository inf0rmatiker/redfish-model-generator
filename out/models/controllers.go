/* -----------------------------------------------------------------
* controllers.go -
*
* DMTF Redfish Controllers resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A network controller ASIC that makes up part of a network adapter.
type Controllers struct {
	// The capabilities of this controller.
	ControllerCapabilities ControllerCapabilities `json:"ControllerCapabilities,omitempty"`

	// The version of the user-facing firmware package.
	FirmwarePackageVersion string `json:"FirmwarePackageVersion,omitempty"`

	// The durable names for the network adapter controller.
	Identifiers array `json:"Identifiers,omitempty"`

	// The links to other resources that are related to this resource.
	Links ControllerLinks `json:"Links,omitempty"`

	// The location of the network adapter controller.
	Location Location `json:"Location,omitempty"`

	// The PCIe interface details for this controller.
	PCIeInterface PCIeInterface `json:"PCIeInterface,omitempty"`

}
