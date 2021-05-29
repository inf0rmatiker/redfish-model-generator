/* -----------------------------------------------------------------
* job_service.go -
*
* DMTF Redfish JobService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The JobService schema contains properties for scheduling and execution of operations, represents the properties for the job service itself, and has links to jobs managed by the job service.
type JobService struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The current date and time setting for the job service.
	DateTime string `json:"DateTime,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to the jobs collection.
	Jobs JobCollection `json:"Jobs,omitempty"`

	// The link to a log service that the job service uses.  This service can be a dedicated log service or a pointer a log service under another resource, such as a manager.
	Log LogService `json:"Log,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The supported capabilities of this job service implementation.
	ServiceCapabilities JobServiceCapabilities `json:"ServiceCapabilities,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
