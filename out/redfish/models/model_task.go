/* -----------------------------------------------------------------
* task.go -
*
* DMTF Redfish Task resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Task - This resource contains information about a specific Task scheduled by or being executed by a Redfish service's Task Service.  Tasks are used to represent operations that take more time than a client typically wants to wait.
// Modeled after DMTF Redfish schema Task
type Task struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The date-time stamp that the task was last completed.
	EndTime string `json:"EndTime,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// This is an array of messages associated with the task.
	Messages []Message `json:"Messages,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The date-time stamp that the task was last started.
	StartTime string `json:"StartTime,omitempty"`

	// The state of the task.
	TaskState TaskState `json:"TaskState,omitempty"`

	// This is the completion status of the task.
	TaskStatus map[string]interface{} `json:"TaskStatus,omitempty"`

}
