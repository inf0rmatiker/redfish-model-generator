/* -----------------------------------------------------------------
* expand.go -
*
* DMTF Redfish Expand resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Expand - Contains information about the use of $expand in the service.
// Modeled after DMTF Redfish schema Expand
type Expand struct {
	// This indicates whether the expand support of asterisk (expand all entries) is supported.
	ExpandAll bool `json:"ExpandAll,omitempty"`

	// This indicates whether the expand support of the $levels qualifier is supported by the service.
	Levels bool `json:"Levels,omitempty"`

	// This indicates whether the expand support of tilde (expand only entries in the Links section) is supported.
	Links bool `json:"Links,omitempty"`

	// This indicates the maximum number value of the $levels qualifier in expand operations.
	MaxLevels float64 `json:"MaxLevels,omitempty"`

	// This indicates whether the expand support of period (expand only entries not in the Links section) is supported.
	NoLinks bool `json:"NoLinks,omitempty"`

}
