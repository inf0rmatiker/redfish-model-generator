/* -----------------------------------------------------------------
* memory_chunk_info.go -
*
* DMTF Redfish MemoryChunkInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The combination of permissions and memory chunk information.
type MemoryChunkInfo struct {
	// Supported IO access capabilities.
	AccessCapabilities array `json:"AccessCapabilities,omitempty"`

	// The access state for this connection.
	AccessState AccessState `json:"AccessState,omitempty"`

	// The specified memory chunk.
	MemoryChunk MemoryChunks `json:"MemoryChunk,omitempty"`

}
