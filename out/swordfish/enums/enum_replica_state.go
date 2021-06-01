/* -----------------------------------------------------------------
* enum_replica_state.go -
*
* SNIA Swordfish ReplicaState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaState string

const (
	//  The link to enable replication is established and source/replica elements are associated, but the data flow has not started.
	ReplicaState_INITIALIZED ReplicaState = "Initialized"

	// Not all the source element data has been copied to the target element.
	ReplicaState_UNSYNCHRONIZED ReplicaState = "Unsynchronized"

	// For the Mirror, Snapshot, or Clone replication, the target represents a copy of the source.
	ReplicaState_SYNCHRONIZED ReplicaState = "Synchronized"

	// The relationship is non-functional due to errors in the source, the target, the path between the two or space constraints.
	ReplicaState_BROKEN ReplicaState = "Broken"

	// Target is split from the source.
	ReplicaState_FRACTURED ReplicaState = "Fractured"

	// The target element was gracefully (or systematically) split from its source element -- consistency is guaranteed.
	ReplicaState_SPLIT ReplicaState = "Split"

	// Data flow has stopped, writes to source element will not be sent to target element.
	ReplicaState_INACTIVE ReplicaState = "Inactive"

	// Data flow between the source and target elements has stopped. Writes to source element are held until the relationship is Resumed.
	ReplicaState_SUSPENDED ReplicaState = "Suspended"

	// Reads and writes are sent to the target element. Source element is not reachable.
	ReplicaState_FAILEDOVER ReplicaState = "Failedover"

	// Initialization is completed, however, the data flow has not started.
	ReplicaState_PREPARED ReplicaState = "Prepared"

	// The copy operation is aborted with the Abort operation. Use the Resync Replica operation to restart the copy operation.
	ReplicaState_ABORTED ReplicaState = "Aborted"

	// The target has been modified and is no longer synchronized with the source element or the point-in-time view.
	ReplicaState_SKEWED ReplicaState = "Skewed"

	// Applies to the ReplicaState of GroupSynchronized. It indicates the StorageSynchronized relationships of the elements in the groups have different ReplicaState values.
	ReplicaState_MIXED ReplicaState = "Mixed"

	// State of replication relationship can not be determined, for example, due to a connection problem.
	ReplicaState_PARTITIONED ReplicaState = "Partitioned"

	// The array is unable to determine the state of the replication relationship, for example, after the connection is restored; however, either source or target elements have an unknown status.
	ReplicaState_INVALID ReplicaState = "Invalid"

	// It indicates the source element was restored from the target element.
	ReplicaState_RESTORED ReplicaState = "Restored"

)
