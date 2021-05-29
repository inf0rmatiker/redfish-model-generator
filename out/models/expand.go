/* -----------------------------------------------------------------
* expand.go -
*
* DMTF Redfish Expand resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The information about the use of $expand in the service.
type Expand struct {
	// An indication of whether the service supports the asterisk (`*`) option of the $expand query parameter.
	ExpandAll bool `json:"ExpandAll,omitempty"`

	// An indication of whether the service supports the $levels option of the $expand query parameter.
	Levels bool `json:"Levels,omitempty"`

	// An indication of whether this service supports the tilde (`~`) option of the $expand query parameter.
	Links bool `json:"Links,omitempty"`

	// The maximum $levels option value in the $expand query parameter.
	MaxLevels int `json:"MaxLevels,omitempty"`

	// An indication of whether the service supports the period (`.`) option of the $expand query parameter.
	NoLinks bool `json:"NoLinks,omitempty"`

}
