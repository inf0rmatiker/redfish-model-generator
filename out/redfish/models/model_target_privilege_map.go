/* -----------------------------------------------------------------
* target_privilege_map.go -
*
* DMTF Redfish Target_PrivilegeMap resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Target_PrivilegeMap - This type describes a mapping between one or more targets and the HTTP operations associated with them.
// Modeled after DMTF Redfish schema Target_PrivilegeMap
type Target_PrivilegeMap struct {
	// List mapping between HTTP operation and privilege needed to perform operation.
	OperationMap OperationMap `json:"OperationMap,omitempty"`

	// Indicates the URI or Entity.
	Targets []string `json:"Targets,omitempty"`

}