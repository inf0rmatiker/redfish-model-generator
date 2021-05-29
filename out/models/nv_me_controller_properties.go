/* -----------------------------------------------------------------
* nv_me_controller_properties.go -
*
* DMTF Redfish NVMeControllerProperties resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NVMe related properties for a storage controller.
type NVMeControllerProperties struct {
	// The ANA characteristics and volume information.
	ANACharacteristics array `json:"ANACharacteristics,omitempty"`

	// The type of NVMe controller.
	ControllerType NVMeControllerType `json:"ControllerType,omitempty"`

	// The maximum individual queue size that an NVMe IO controller supports.
	MaxQueueSize int `json:"MaxQueueSize,omitempty"`

	// The NVMe controller attributes.
	NVMeControllerAttributes NVMeControllerAttributes `json:"NVMeControllerAttributes,omitempty"`

	// The NVMe SMART Critical Warnings for this storage controller.  This property contains possible triggers for the predictive drive failure warning for the corresponding drive.
	NVMeSMARTCriticalWarnings NVMeSMARTCriticalWarnings `json:"NVMeSMARTCriticalWarnings,omitempty"`

	// The version of the NVMe Base Specification supported.
	NVMeVersion string `json:"NVMeVersion,omitempty"`

}
