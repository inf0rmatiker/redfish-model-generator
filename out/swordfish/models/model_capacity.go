/* -----------------------------------------------------------------
* capacity.go -
*
* SNIA Swordfish Capacity resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Capacity - This is the schema definition for the Capacity of a device. It represents the properties for capacity for any data store.
// Modeled after SNIA Swordfish schema Capacity
type Capacity struct {
	// The capacity information relating to the user data.
	Data CapacityInfo `json:"Data,omitempty"`

	// Marks that the capacity is not necessarily fully allocated.
	IsThinProvisioned bool `json:"IsThinProvisioned,omitempty"`

	// The capacity information relating to  metadata.
	Metadata CapacityInfo `json:"Metadata,omitempty"`

	// The capacity information relating to snapshot or backup data.
	Snapshot CapacityInfo `json:"Snapshot,omitempty"`

}
