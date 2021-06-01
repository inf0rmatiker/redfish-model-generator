/* -----------------------------------------------------------------
* operation_map.go -
*
* DMTF Redfish OperationMap resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// OperationMap - Used for describing the specific privileges for a set of HTTP operations.
// Modeled after DMTF Redfish schema OperationMap
type OperationMap struct {
	// Indicates privilege required for HTTP DELETE operation.
	DELETE []OperationPrivilege `json:"DELETE,omitempty"`

	// Indicates privilege required for HTTP GET operation.
	GET []OperationPrivilege `json:"GET,omitempty"`

	// Indicates privilege required for HTTP HEAD operation.
	HEAD []OperationPrivilege `json:"HEAD,omitempty"`

	// Indicates privilege required for HTTP PATCH operation.
	PATCH []OperationPrivilege `json:"PATCH,omitempty"`

	// Indicates privilege required for HTTP POST operation.
	POST []OperationPrivilege `json:"POST,omitempty"`

	// Indicates privilege required for HTTP PUT operation.
	PUT []OperationPrivilege `json:"PUT,omitempty"`

}
