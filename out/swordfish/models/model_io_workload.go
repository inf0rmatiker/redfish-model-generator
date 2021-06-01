/* -----------------------------------------------------------------
* io_workload.go -
*
* SNIA Swordfish IOWorkload resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IOWorkload - An IO Workload description.
// Modeled after SNIA Swordfish schema IOWorkload
type IOWorkload struct {
	// An array of IO workload component descriptions.
	Components []IOWorkloadComponent `json:"Components,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

}
