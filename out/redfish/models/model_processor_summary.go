/* -----------------------------------------------------------------
* processor_summary.go -
*
* DMTF Redfish ProcessorSummary resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ProcessorSummary - The central processors of the system in general detail.
// Modeled after DMTF Redfish schema ProcessorSummary
type ProcessorSummary struct {
	// The number of physical processors in the system.
	Count int `json:"Count,omitempty"`

	// The processor model for the primary or majority of processors in this system.
	Model string `json:"Model,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
