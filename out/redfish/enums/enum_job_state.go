/* -----------------------------------------------------------------
* enum_job_state.go -
*
* DMTF Redfish JobState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type JobState string

const (
	// A new job.
	JobState_NEW JobState = "New"

	// Job is starting.
	JobState_STARTING JobState = "Starting"

	// Job is running normally.
	JobState_RUNNING JobState = "Running"

	// Job has been suspended.
	JobState_SUSPENDED JobState = "Suspended"

	// Job has been interrupted.
	JobState_INTERRUPTED JobState = "Interrupted"

	// Job is pending and has not started.
	JobState_PENDING JobState = "Pending"

	// Job is in the process of stopping.
	JobState_STOPPING JobState = "Stopping"

	// Job was completed.
	JobState_COMPLETED JobState = "Completed"

	// Job was cancelled.
	JobState_CANCELLED JobState = "Cancelled"

	// Job has stopped due to an exception condition.
	JobState_EXCEPTION JobState = "Exception"

	// Job is running as a service.
	JobState_SERVICE JobState = "Service"

	// Job is waiting for user intervention.
	JobState_USER_INTERVENTION JobState = "UserIntervention"

	// Job is to resume operation.
	JobState_CONTINUE JobState = "Continue"

)
