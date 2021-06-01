/* -----------------------------------------------------------------
* composition_status.go -
*
* DMTF Redfish CompositionStatus resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CompositionStatus - Composition status of the Resource Block.
// Modeled after DMTF Redfish schema CompositionStatus
type CompositionStatus struct {
	// This property represents the current state of the Resource Block from a composition perspective.
	CompositionState CompositionState `json:"CompositionState"`

	// The maximum number of compositions in which this Resource Block is capable of participating simultaneously.
	MaxCompositions int `json:"MaxCompositions,omitempty"`

	// The number of compositions in which this Resource Block is currently participating.
	NumberOfCompositions int `json:"NumberOfCompositions,omitempty"`

	// This represents if the Resource Block is reserved by any client.
	Reserved bool `json:"Reserved,omitempty"`

	// Indicates if this Resource Block is capable of participating in multiple compositions simultaneously.
	SharingCapable bool `json:"SharingCapable,omitempty"`

	// Indicates if this Resource Block is allowed to participate in multiple compositions simultaneously.
	SharingEnabled bool `json:"SharingEnabled,omitempty"`

}
