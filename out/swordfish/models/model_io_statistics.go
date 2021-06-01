/* -----------------------------------------------------------------
* io_statistics.go -
*
* SNIA Swordfish IOStatistics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IOStatistics - The properties of this type represent IO statistics.
// Modeled after SNIA Swordfish schema IOStatistics
type IOStatistics struct {
	// The time that the resource is busy processing write requests.
	NonIORequestTime string `json:"NonIORequestTime,omitempty"`

	// Count of non IO requests.
	NonIORequests int `json:"NonIORequests,omitempty"`

	// Count of read IO requests satisfied from memory.
	ReadHitIORequests int `json:"ReadHitIORequests,omitempty"`

	// Number of kibibytes read.
	ReadIOKiBytes int `json:"ReadIOKiBytes,omitempty"`

	// The time that the resource is busy processing read requests.
	ReadIORequestTime string `json:"ReadIORequestTime,omitempty"`

	// Count of read IO requests.
	ReadIORequests int `json:"ReadIORequests,omitempty"`

	// Count of write IO requests coalesced into memory.
	WriteHitIORequests int `json:"WriteHitIORequests,omitempty"`

	// Number of kibibytes written.
	WriteIOKiBytes int `json:"WriteIOKiBytes,omitempty"`

	// The time that the resource is busy processing write requests.
	WriteIORequestTime string `json:"WriteIORequestTime,omitempty"`

	// Count of write IO requests.
	WriteIORequests int `json:"WriteIORequests,omitempty"`

}
