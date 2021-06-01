/* -----------------------------------------------------------------
* enum_replica_progress_status.go -
*
* SNIA Swordfish ReplicaProgressStatus enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaProgressStatus string

const (
	// The request is completed. Data flow is idle.
	ReplicaProgressStatus_COMPLETED ReplicaProgressStatus = "Completed"

	// Indicates that the data flow is inactive, suspended or quiesced.
	ReplicaProgressStatus_DORMANT ReplicaProgressStatus = "Dormant"

	// In the process of establishing source/replica relationship and the data flow has not started.
	ReplicaProgressStatus_INITIALIZING ReplicaProgressStatus = "Initializing"

	// Preparation in progress.
	ReplicaProgressStatus_PREPARING ReplicaProgressStatus = "Preparing"

	// Sync in progress.
	ReplicaProgressStatus_SYNCHRONIZING ReplicaProgressStatus = "Synchronizing"

	// Resync in progress.
	ReplicaProgressStatus_RESYNCING ReplicaProgressStatus = "Resyncing"

	// Restore in progress.
	ReplicaProgressStatus_RESTORING ReplicaProgressStatus = "Restoring"

	// Fracture in progress.
	ReplicaProgressStatus_FRACTURING ReplicaProgressStatus = "Fracturing"

	// Split in progress.
	ReplicaProgressStatus_SPLITTING ReplicaProgressStatus = "Splitting"

	// In the process of switching source and target.
	ReplicaProgressStatus_FAILING_OVER ReplicaProgressStatus = "FailingOver"

	// Undoing the result of failover.
	ReplicaProgressStatus_FAILING_BACK ReplicaProgressStatus = "FailingBack"

	// Detach in progress.
	ReplicaProgressStatus_DETACHING ReplicaProgressStatus = "Detaching"

	// Abort in progress.
	ReplicaProgressStatus_ABORTING ReplicaProgressStatus = "Aborting"

	// Applies to groups with element pairs with different statuses. Generally, the individual statuses need to be examined.
	ReplicaProgressStatus_MIXED ReplicaProgressStatus = "Mixed"

	// The copy operation is in the process of being suspended.
	ReplicaProgressStatus_SUSPENDING ReplicaProgressStatus = "Suspending"

	// The requested operation has completed, however, the synchronization relationship needs to be fractured before further copy operations can be issued.
	ReplicaProgressStatus_REQUIRES_FRACTURE ReplicaProgressStatus = "RequiresFracture"

	// The requested operation has completed, however, the synchronization relationship needs to be resynced before further copy operations can be issued.
	ReplicaProgressStatus_REQUIRES_RESYNC ReplicaProgressStatus = "RequiresResync"

	// The requested operation has completed, however, the synchronization relationship needs to be activated before further copy operations can be issued.
	ReplicaProgressStatus_REQUIRES_ACTIVATE ReplicaProgressStatus = "RequiresActivate"

	// The flow of data has stopped momentarily due to limited bandwidth or a busy system.
	ReplicaProgressStatus_PENDING ReplicaProgressStatus = "Pending"

	// The requested operation has completed, however, the synchronization relationship needs to be detached before further copy operations can be issued.
	ReplicaProgressStatus_REQUIRES_DETACH ReplicaProgressStatus = "RequiresDetach"

	// The relationship is in the process of terminating.
	ReplicaProgressStatus_TERMINATING ReplicaProgressStatus = "Terminating"

	// The requested operation has completed, however, the synchronization relationship needs to be split before further copy operations can be issued.
	ReplicaProgressStatus_REQUIRES_SPLIT ReplicaProgressStatus = "RequiresSplit"

	// The requested operation has completed, however, the synchronization relationship needs to be resumed before further copy operations can be issued.
	ReplicaProgressStatus_REQUIRES_RESUME ReplicaProgressStatus = "RequiresResume"

)
