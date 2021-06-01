/* -----------------------------------------------------------------
* role_mapping.go -
*
* DMTF Redfish RoleMapping resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// RoleMapping - The mapping rules that are used to convert the external account providers account information to the local Redfish role.
// Modeled after DMTF Redfish schema RoleMapping
type RoleMapping struct {
	// The name of the local Redfish role to which to map the remote user or group.
	LocalRole string `json:"LocalRole,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The name of the remote group, or the remote role in the case of a Redfish service, that maps to the local Redfish role to which this entity links.
	RemoteGroup string `json:"RemoteGroup,omitempty"`

	// The name of the remote user that maps to the local Redfish role to which this entity links.
	RemoteUser string `json:"RemoteUser,omitempty"`

}
