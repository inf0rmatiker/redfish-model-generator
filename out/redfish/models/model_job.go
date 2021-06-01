/* -----------------------------------------------------------------
* job.go -
*
* DMTF Redfish Job resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Job - The Job schema contains information about a job that a Redfish job service schedules or executes.  Clients create jobs to describe a series of operations that occur at periodic intervals.
// Modeled after DMTF Redfish schema Job
type Job struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The person or program that created this job entry.
	CreatedBy string `json:"CreatedBy,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The date and time when the job was completed.
	EndTime string `json:"EndTime,omitempty"`

	// An indication of whether the contents of the payload should be hidden from view after the job has been created.  If `true`, responses do not return the payload.  If `false`, responses return the payload.  If this property is not present when the job is created, the default is `false`.
	HidePayload bool `json:"HidePayload,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The state of the job.
	JobState JobState `json:"JobState,omitempty"`

	// The status of the job.
	JobStatus map[string]interface{} `json:"JobStatus,omitempty"`

	// The maximum amount of time the job is allowed to execute.
	MaxExecutionTime string `json:"MaxExecutionTime,omitempty"`

	// An array of messages associated with the job.
	Messages []Message `json:"Messages,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The HTTP and JSON payload details for this job.
	Payload Payload `json:"Payload,omitempty"`

	// The completion percentage of this job.
	PercentComplete int `json:"PercentComplete,omitempty"`

	// The schedule settings for this job.
	Schedule map[string]interface{} `json:"Schedule,omitempty"`

	// The date and time when the job was started or is scheduled to start.
	StartTime string `json:"StartTime,omitempty"`

	// The serialized execution order of the job steps.
	StepOrder []s `json:"StepOrder,omitempty"`

	// The link to a collection of steps for this job.
	Steps map[string]interface{} `json:"Steps,omitempty"`

}
