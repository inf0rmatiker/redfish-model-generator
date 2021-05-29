/* -----------------------------------------------------------------
* task.go -
*
* DMTF Redfish Task resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Task schema contains information about a task that the Redfish Task Service schedules or executes.  Tasks represent operations that take more time than a client typically wants to wait.
type Task struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The date and time when the task was completed.  This property will only appear when the task is complete.
	EndTime string `json:"EndTime,omitempty"`

	// An indication of whether the contents of the payload are hidden from view after the task has been created.  If `true`, responses do not return the payload.  If `false`, responses return the payload.  If this property is not present when the task is created, the default is `false`.
	HidePayload bool `json:"HidePayload,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An array of messages associated with the task.
	Messages array `json:"Messages,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The HTTP and JSON payload details for this task, unless they are hidden from view by the service.
	Payload Payload `json:"Payload,omitempty"`

	// The completion percentage of this task.
	PercentComplete int `json:"PercentComplete,omitempty"`

	// The date and time when the task was started.
	StartTime string `json:"StartTime,omitempty"`

	// The link to a collection of sub-tasks for this task.
	SubTasks TaskCollection `json:"SubTasks,omitempty"`

	// The URI of the Task Monitor for this task.
	TaskMonitor string `json:"TaskMonitor,omitempty"`

	// The state of the task.
	TaskState TaskState `json:"TaskState,omitempty"`

	// The completion status of the task.
	TaskStatus Health `json:"TaskStatus,omitempty"`

}
