/* -----------------------------------------------------------------
* task_service.go -
*
* DMTF Redfish TaskService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The TaskService schema describes a task service that enables management of long-duration operations, includes the properties for the task service itself, and has links to the resource collection of tasks.
type TaskService struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The overwrite policy for completed tasks.  This property indicates if the task service overwrites completed task information.
	CompletedTaskOverWritePolicy OverWritePolicy `json:"CompletedTaskOverWritePolicy,omitempty"`

	// The current date and time, with UTC offset, setting that the task service uses.
	DateTime string `json:"DateTime,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether a task state change sends an event.
	LifeCycleEventOnTaskStateChange bool `json:"LifeCycleEventOnTaskStateChange,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The number of minutes after which a completed task is deleted by the service.
	TaskAutoDeleteTimeoutMinutes int `json:"TaskAutoDeleteTimeoutMinutes,omitempty"`

	// The links to the collection of tasks.
	Tasks TaskCollection `json:"Tasks,omitempty"`

}
