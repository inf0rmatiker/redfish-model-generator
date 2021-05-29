/* -----------------------------------------------------------------
* calculation_params_type.go -
*
* DMTF Redfish CalculationParamsType resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The usage of the parameter in the calculation.
type CalculationParamsType struct {
	// The link to a metric property that stores the result of the calculation.  If the link has wildcards, the wildcards are substituted as specified in the Wildcards array property.
	ResultMetric string `json:"ResultMetric,omitempty"`

	// The metric property used as the input into the calculation.  If the link has wildcards, the wildcards are substituted as specified in the Wildcards array property.
	SourceMetric string `json:"SourceMetric,omitempty"`

}
