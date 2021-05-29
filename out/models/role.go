/* -----------------------------------------------------------------
* role.go -
*
* DMTF Redfish Role resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Role schema contains a Redfish role to use in conjunction with a manager account.
type Role struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An equivalent role to use when this role is restricted.
	AlternateRoleId string `json:"AlternateRoleId,omitempty"`

	// The Redfish privileges for this role.
	AssignedPrivileges array `json:"AssignedPrivileges,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether the role is predefined by Redfish or an OEM rather than a client-defined role.
	IsPredefined bool `json:"IsPredefined,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The OEM privileges for this role.
	OemPrivileges []string `json:"OemPrivileges,omitempty"`

	// An indication of whether use of the role is restricted.
	Restricted bool `json:"Restricted,omitempty"`

	// The name of the role.
	RoleId string `json:"RoleId,omitempty"`

}
