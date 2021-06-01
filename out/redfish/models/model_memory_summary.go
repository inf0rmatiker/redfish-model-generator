/* -----------------------------------------------------------------
* memory_summary.go -
*
* DMTF Redfish MemorySummary resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MemorySummary - The memory of the system in general detail.
// Modeled after DMTF Redfish schema MemorySummary
type MemorySummary struct {
	// The ability and type of memory mirroring that this computer system supports.
	MemoryMirroring MemoryMirroring `json:"MemoryMirroring,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The total configured operating system-accessible memory (RAM), measured in GiB.
	TotalSystemMemoryGiB float64 `json:"TotalSystemMemoryGiB,omitempty"`

	// The total configured, system-accessible persistent memory, measured in GiB.
	TotalSystemPersistentMemoryGiB float64 `json:"TotalSystemPersistentMemoryGiB,omitempty"`

}
