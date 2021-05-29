/* -----------------------------------------------------------------
* job_service_capabilities.go -
*
* DMTF Redfish JobServiceCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The supported capabilities of this job service implementation.
type JobServiceCapabilities struct {
	// The maximum number of jobs supported.
	MaxJobs int `json:"MaxJobs,omitempty"`

	// The maximum number of job steps supported.
	MaxSteps int `json:"MaxSteps,omitempty"`

	// An indication of whether scheduling of jobs is supported.
	Scheduling bool `json:"Scheduling,omitempty"`

}
