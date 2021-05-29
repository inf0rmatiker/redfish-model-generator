/* -----------------------------------------------------------------
* privilege_registry.go -
*
* DMTF Redfish PrivilegeRegistry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The PrivilegeRegistry schema describes the operation-to-privilege mappings.
type PrivilegeRegistry struct {
	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The mappings between entities and the relevant privileges that access those entities.
	Mappings array `json:"Mappings,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The set of OEM privileges used in this mapping.
	OEMPrivilegesUsed []string `json:"OEMPrivilegesUsed,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The set of Redfish standard privileges used in this mapping.
	PrivilegesUsed array `json:"PrivilegesUsed,omitempty"`

}
