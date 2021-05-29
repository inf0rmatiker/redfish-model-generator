/* -----------------------------------------------------------------
* assembly.go -
*
* DMTF Redfish Assembly resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Assembly schema defines an assembly.  Assembly information contains details about a device, such as part number, serial number, manufacturer, and production date.  It also provides access to the original data for the assembly.
type Assembly struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The assembly records.
	Assemblies array `json:"Assemblies,omitempty"`

	Assemblies@odata.count count `json:"Assemblies@odata.count,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
