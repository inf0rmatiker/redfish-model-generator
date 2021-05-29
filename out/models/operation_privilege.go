/* -----------------------------------------------------------------
* operation_privilege.go -
*
* DMTF Redfish OperationPrivilege resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The privileges for a specific HTTP operation.
type OperationPrivilege struct {
	// An array of privileges that are required to complete a specific HTTP operation on a Resource.
	Privilege []string `json:"Privilege,omitempty"`

}
