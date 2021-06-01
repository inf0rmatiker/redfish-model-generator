/* -----------------------------------------------------------------
* operations.go -
*
* DMTF Redfish Operations resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Operations - An operation currently running on this resource.
// Modeled after DMTF Redfish schema Operations
type Operations struct {
	// The link to the task associated with the operation, if any.
	AssociatedTask map[string]interface{} `json:"AssociatedTask,omitempty"`

	// The name of the operation.
	OperationName string `json:"OperationName,omitempty"`

	// The percentage of the operation that has been completed.
	PercentageComplete int `json:"PercentageComplete,omitempty"`

}
