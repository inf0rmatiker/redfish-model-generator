/* -----------------------------------------------------------------
* acceleration_function.go -
*
* DMTF Redfish AccelerationFunction resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The AccelerationFunction schema describes an acceleration function that a processor implements.  This can include functions such as audio processing, compression, encryption, packet inspection, packet switching, scheduling, or video processing.
type AccelerationFunction struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The acceleration function type.
	AccelerationFunctionType AccelerationFunctionType `json:"AccelerationFunctionType,omitempty"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An array of the reconfiguration slot identifiers of the FPGA that this acceleration function occupies.
	FpgaReconfigurationSlots []string `json:"FpgaReconfigurationSlots,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links,omitempty"`

	// The acceleration function code manufacturer.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The acceleration function power consumption, in watts.
	PowerWatts int `json:"PowerWatts,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

	// The UUID for this acceleration function.
	UUID UUID `json:"UUID,omitempty"`

	// The acceleration function version.
	Version string `json:"Version,omitempty"`

}
