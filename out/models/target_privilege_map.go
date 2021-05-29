/* -----------------------------------------------------------------
* target_privilege_map.go -
*
* DMTF Redfish Target_PrivilegeMap resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes a mapping between one or more targets and the HTTP operations associated with them.
type Target_PrivilegeMap struct {
	// The mapping between the HTTP operation and the privilege required to complete the operation.
	OperationMap OperationMap `json:"OperationMap,omitempty"`

	// The set of URIs, Resource types, or properties.
	Targets []string `json:"Targets,omitempty"`

}
