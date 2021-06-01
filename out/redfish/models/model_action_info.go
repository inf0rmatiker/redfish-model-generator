/* -----------------------------------------------------------------
* action_info.go -
*
* DMTF Redfish ActionInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ActionInfo - The ActionInfo schema describes the parameters and other information necessary to perform a Redfish Action on a particular Action target. Parameter support can differ between vendors and even between instances of a resource. This data can be used to ensure Action requests from applications contain supported parameters.
// Modeled after DMTF Redfish schema ActionInfo
type ActionInfo struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The parameters associated with the specified Redfish Action.
	Parameters []Parameters `json:"Parameters,omitempty"`

}
