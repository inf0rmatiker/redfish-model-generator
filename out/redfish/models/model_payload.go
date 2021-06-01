/* -----------------------------------------------------------------
* payload.go -
*
* DMTF Redfish Payload resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Payload - The HTTP and JSON payload details for this job.
// Modeled after DMTF Redfish schema Payload
type Payload struct {
	// An array of HTTP headers in this job.
	HttpHeaders []s `json:"HttpHeaders,omitempty"`

	// The HTTP operation that executes this job.
	HttpOperation string `json:"HttpOperation,omitempty"`

	// The JSON payload to use in the execution of this job.
	JsonBody string `json:"JsonBody,omitempty"`

	// The link to the target for this job.
	TargetUri string `json:"TargetUri,omitempty"`

}
