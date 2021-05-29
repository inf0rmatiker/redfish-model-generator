/* -----------------------------------------------------------------
* operations.go -
*
* DMTF Redfish Operations resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// An operation currently running on this resource.
type Operations struct {
	// The link to the task associated with the operation, if any.
	AssociatedTask Task `json:"AssociatedTask,omitempty"`

	// The name of the operation.
	OperationName string `json:"OperationName,omitempty"`

	// The percentage of the operation that has been completed.
	PercentageComplete int `json:"PercentageComplete,omitempty"`

}
