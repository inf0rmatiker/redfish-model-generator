/* -----------------------------------------------------------------
* job_service.go -
*
* DMTF Redfish JobService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// JobService - This is the schema definition for the Job Service.  It represents the properties for the service itself and has links to the actual list of tasks.
// Modeled after DMTF Redfish schema JobService
type JobService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The current DateTime (with offset) setting that the job service is using.
	DateTime string `json:"DateTime,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// References to the Jobs collection.
	Jobs map[string]interface{} `json:"Jobs,omitempty"`

	// This is a reference to a Log Service used by the Job Service.
	Log map[string]interface{} `json:"Log,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This object describes the supported capabilities of this Job Service implementation.
	ServiceCapabilities JobServiceCapabilities `json:"ServiceCapabilities,omitempty"`

	// This indicates whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

}
