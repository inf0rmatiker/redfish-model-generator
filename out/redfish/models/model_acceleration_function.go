/* -----------------------------------------------------------------
* acceleration_function.go -
*
* DMTF Redfish AccelerationFunction resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// AccelerationFunction - The AccelerationFunction schema defines the accelerator implemented in a Processor device.
// Modeled after DMTF Redfish schema AccelerationFunction
type AccelerationFunction struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The type of acceleration function.
	AccelerationFunctionType AccelerationFunctionType `json:"AccelerationFunctionType,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An array of the reconfiguration slot identifiers for an FPGA.
	FpgaReconfigurationSlots []s `json:"FpgaReconfigurationSlots,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Contains references to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The acceleration function code manufacturer.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The acceleration function power consumption.
	PowerWatts int `json:"PowerWatts,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The universal unique identifier (UUID) for this acceleration function.
	UUID map[string]interface{} `json:"UUID,omitempty"`

	// The acceleration function version.
	Version string `json:"Version,omitempty"`

}
