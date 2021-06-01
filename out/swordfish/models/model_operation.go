/* -----------------------------------------------------------------
* operation.go -
*
* SNIA Swordfish Operation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// A reference to the task associated with the operation if any.
	AssociatedFeaturesRegistry map[string]interface{} `json:"AssociatedFeaturesRegistry,omitempty"`

	// The name of the operation.
	OperationName string `json:"OperationName,omitempty"`

	// The percentage of the operation that has been completed.
	PercentageComplete int `json:"PercentageComplete,omitempty"`

}
