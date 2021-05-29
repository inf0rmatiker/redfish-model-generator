/* -----------------------------------------------------------------
* action_info.go -
*
* DMTF Redfish ActionInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The ActionInfo schema defines the supported parameters and other information for a Redfish action.  Supported parameters can differ among vendors and even among Resource instances.  This data can ensure that action requests from applications contain supported parameters.
type ActionInfo struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The list of parameters included in the specified Redfish action.
	Parameters array `json:"Parameters,omitempty"`

}
