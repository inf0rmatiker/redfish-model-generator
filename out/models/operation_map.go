/* -----------------------------------------------------------------
* operation_map.go -
*
* DMTF Redfish OperationMap resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The specific privileges required to complete a set of HTTP operations.
type OperationMap struct {
	// The privilege required to complete an HTTP DELETE operation.
	DELETE array `json:"DELETE,omitempty"`

	// The privilege required to complete an HTTP GET operation.
	GET array `json:"GET,omitempty"`

	// The privilege required to complete an HTTP HEAD operation.
	HEAD array `json:"HEAD,omitempty"`

	// The privilege required to complete an HTTP PATCH operation.
	PATCH array `json:"PATCH,omitempty"`

	// The privilege required to complete an HTTP POST operation.
	POST array `json:"POST,omitempty"`

	// The privilege required to complete an HTTP PUT operation.
	PUT array `json:"PUT,omitempty"`

}
