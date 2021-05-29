/* -----------------------------------------------------------------
* memory_location.go -
*
* DMTF Redfish MemoryLocation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Memory connection information to sockets and memory controllers.
type MemoryLocation struct {
	// The channel number to which the memory device is connected.
	Channel int `json:"Channel,omitempty"`

	// The memory controller number to which the memory device is connected.
	MemoryController int `json:"MemoryController,omitempty"`

	// The slot number to which the memory device is connected.
	Slot int `json:"Slot,omitempty"`

	// The socket number to which the memory device is connected.
	Socket int `json:"Socket,omitempty"`

}
