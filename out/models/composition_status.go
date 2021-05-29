/* -----------------------------------------------------------------
* composition_status.go -
*
* DMTF Redfish CompositionStatus resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Composition status of the resource block.
type CompositionStatus struct {
	// The current state of the resource block from a composition perspective.
	CompositionState CompositionState `json:"CompositionState"`

	// The maximum number of compositions in which this resource block can participate simultaneously.
	MaxCompositions int `json:"MaxCompositions,omitempty"`

	// The number of compositions in which this resource block is currently participating.
	NumberOfCompositions int `json:"NumberOfCompositions,omitempty"`

	// An indication of whether any client has reserved the resource block.
	Reserved bool `json:"Reserved,omitempty"`

	// An indication of whether this resource block can participate in multiple compositions simultaneously.
	SharingCapable bool `json:"SharingCapable,omitempty"`

	// An indication of whether this resource block is allowed to participate in multiple compositions simultaneously.
	SharingEnabled bool `json:"SharingEnabled,omitempty"`

}
