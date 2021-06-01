/* -----------------------------------------------------------------
* privilege_registry.go -
*
* DMTF Redfish PrivilegeRegistry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PrivilegeRegistry - This is the schema definition for Operation to Privilege mapping.
// Modeled after DMTF Redfish schema PrivilegeRegistry
type PrivilegeRegistry struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// This property describes the mappings between entities and the relevant privileges used to access them.
	Mappings []Mapping `json:"Mappings,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// Lists the set of OEM Priviliges used in building this mapping.
	OEMPrivilegesUsed []s `json:"OEMPrivilegesUsed,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Lists the set of Redfish standard priviliges used in building this mapping.
	PrivilegesUsed []PrivilegeType `json:"PrivilegesUsed,omitempty"`

}
