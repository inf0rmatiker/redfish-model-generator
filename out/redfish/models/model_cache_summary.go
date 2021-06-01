/* -----------------------------------------------------------------
* cache_summary.go -
*
* DMTF Redfish CacheSummary resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CacheSummary - This type describes the cache memory of the storage controller in general detail.
// Modeled after DMTF Redfish schema CacheSummary
type CacheSummary struct {
	// The portion of the cache memory that is persistent, measured in MiB.
	PersistentCacheSizeMiB int `json:"PersistentCacheSizeMiB,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The total configured cache memory, measured in MiB.
	TotalCacheSizeMiB int `json:"TotalCacheSizeMiB"`

}
