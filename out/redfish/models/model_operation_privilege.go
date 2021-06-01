/* -----------------------------------------------------------------
* operation_privilege.go -
*
* DMTF Redfish OperationPrivilege resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// OperationPrivilege - Used for describing the specific privileges for a given type of HTTP operation.
// Modeled after DMTF Redfish schema OperationPrivilege
type OperationPrivilege struct {
	// Lists the privileges that are allowed to perform the given type of HTTP operation on the entity type.
	Privilege []s `json:"Privilege,omitempty"`

}
