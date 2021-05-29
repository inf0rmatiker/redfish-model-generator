/* -----------------------------------------------------------------
* rates.go -
*
* DMTF Redfish Rates resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes the various controller rates used for processes such as volume rebuild or consistency checks.
type Rates struct {
	// The percentage of controller resources used for performing a data consistency check on volumes.
	ConsistencyCheckRatePercent int `json:"ConsistencyCheckRatePercent,omitempty"`

	// The percentage of controller resources used for rebuilding/repairing volumes.
	RebuildRatePercent int `json:"RebuildRatePercent,omitempty"`

	// The percentage of controller resources used for transforming volumes from one configuration to another.
	TransformationRatePercent int `json:"TransformationRatePercent,omitempty"`

}
