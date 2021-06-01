/* -----------------------------------------------------------------
* enum_task_state.go -
*
* DMTF Redfish TaskState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type TaskState string

const (
	// A new task.
	TaskState_NEW TaskState = "New"

	// Task is starting.
	TaskState_STARTING TaskState = "Starting"

	// Task is running normally.
	TaskState_RUNNING TaskState = "Running"

	// Task has been suspended.
	TaskState_SUSPENDED TaskState = "Suspended"

	// Task has been interrupted.
	TaskState_INTERRUPTED TaskState = "Interrupted"

	// Task is pending and has not started.
	TaskState_PENDING TaskState = "Pending"

	// Task is in the process of stopping.
	TaskState_STOPPING TaskState = "Stopping"

	// Task has completed.
	TaskState_COMPLETED TaskState = "Completed"

	// Task was terminated.
	TaskState_KILLED TaskState = "Killed"

	// Task has stopped due to an exception condition.
	TaskState_EXCEPTION TaskState = "Exception"

	// Task is running as a service.
	TaskState_SERVICE TaskState = "Service"

)
