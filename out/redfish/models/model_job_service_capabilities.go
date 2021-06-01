/* -----------------------------------------------------------------
* job_service_capabilities.go -
*
* DMTF Redfish JobServiceCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// JobServiceCapabilities - This object describes the supported capabilities of this Job Service implementation.
// Modeled after DMTF Redfish schema JobServiceCapabilities
type JobServiceCapabilities struct {
	// Maximum number of Jobs supported.
	MaxJobs int `json:"MaxJobs,omitempty"`

	// Maximum number of Job Steps supported.
	MaxSteps int `json:"MaxSteps,omitempty"`

	// Indicates whether scheduling of Jobs is supported.
	Scheduling bool `json:"Scheduling,omitempty"`

}
