/* -----------------------------------------------------------------
* capacity_info.go -
*
* SNIA Swordfish CapacityInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CapacityInfo - The capacity of specific data type in a data store.
// Modeled after SNIA Swordfish schema CapacityInfo
type CapacityInfo struct {
	// The number of bytes currently allocated by the storage system in this data store for this data type.
	AllocatedBytes int `json:"AllocatedBytes,omitempty"`

	// The number of bytes consumed in this data store for this data type.
	ConsumedBytes int `json:"ConsumedBytes,omitempty"`

	// The number of bytes the storage system guarantees can be allocated in this data store for this data type.
	GuaranteedBytes int `json:"GuaranteedBytes,omitempty"`

	// The maximum number of bytes that can be allocated in this data store for this data type.
	ProvisionedBytes int `json:"ProvisionedBytes,omitempty"`

}
