/* -----------------------------------------------------------------
* memory_summary.go -
*
* DMTF Redfish MemorySummary resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The memory of the system in general detail.
type MemorySummary struct {
	// The ability and type of memory mirroring that this computer system supports.
	MemoryMirroring MemoryMirroring `json:"MemoryMirroring,omitempty"`

	// The link to the metrics associated with all memory in this system.
	Metrics MemoryMetrics `json:"Metrics,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The total configured operating system-accessible memory (RAM), measured in GiB.
	TotalSystemMemoryGiB float64 `json:"TotalSystemMemoryGiB,omitempty"`

	// The total configured, system-accessible persistent memory, measured in GiB.
	TotalSystemPersistentMemoryGiB float64 `json:"TotalSystemPersistentMemoryGiB,omitempty"`

}
