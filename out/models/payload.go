/* -----------------------------------------------------------------
* payload.go -
*
* DMTF Redfish Payload resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The HTTP and JSON payload details for this Task.
type Payload struct {
	// An array of HTTP headers that this task includes.
	HttpHeaders []string `json:"HttpHeaders,omitempty"`

	// The HTTP operation to perform to execute this task.
	HttpOperation string `json:"HttpOperation,omitempty"`

	// The JSON payload to use in the execution of this task.
	JsonBody string `json:"JsonBody,omitempty"`

	// The URI of the target for this task.
	TargetUri string `json:"TargetUri,omitempty"`

}
