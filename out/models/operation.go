/* -----------------------------------------------------------------
* operation.go -
*
* DMTF Redfish Operation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Operation struct {
	// A reference to the task associated with the operation if any.
	AssociatedFeaturesRegistry FeaturesRegistry `json:"AssociatedFeaturesRegistry,omitempty"`

	// The name of the operation.
	OperationName string `json:"OperationName,omitempty"`

	// The percentage of the operation that has been completed.
	PercentageComplete int `json:"PercentageComplete,omitempty"`

}
