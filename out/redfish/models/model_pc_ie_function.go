/* -----------------------------------------------------------------
* pc_ie_function.go -
*
* DMTF Redfish PCIeFunction resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PCIeFunction - The schema definition for the PCIeFunction Resource.  It represents the properties of a PCIeFunction attached to a System.
// Modeled after DMTF Redfish schema PCIeFunction
type PCIeFunction struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The Class Code of this PCIe function.
	ClassCode string `json:"ClassCode,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The class for this PCIe function.
	DeviceClass DeviceClass `json:"DeviceClass,omitempty"`

	// The Device ID of this PCIe function.
	DeviceId string `json:"DeviceId,omitempty"`

	// The PCIe Function Number.
	FunctionId int `json:"FunctionId,omitempty"`

	// The type of the PCIe function.
	FunctionType FunctionType `json:"FunctionType,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The Revision ID of this PCIe function.
	RevisionId string `json:"RevisionId,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The Subsystem ID of this PCIe function.
	SubsystemId string `json:"SubsystemId,omitempty"`

	// The Subsystem Vendor ID of this PCIe function.
	SubsystemVendorId string `json:"SubsystemVendorId,omitempty"`

	// The Vendor ID of this PCIe function.
	VendorId string `json:"VendorId,omitempty"`

}
