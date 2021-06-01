/* -----------------------------------------------------------------
* role.go -
*
* DMTF Redfish Role resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Role - The Role schema contains a Redfish Role to use in conjunction with a manager account.
// Modeled after DMTF Redfish schema Role
type Role struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The Redfish privileges for this Role.
	AssignedPrivileges []PrivilegeType `json:"AssignedPrivileges,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether the Role is a Redfish-predefined Role rather than a custom Redfish Role.
	IsPredefined bool `json:"IsPredefined,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The OEM privileges for this Role.
	OemPrivileges []s `json:"OemPrivileges,omitempty"`

}
