/* -----------------------------------------------------------------
* replica_info.go -
*
* SNIA Swordfish ReplicaInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ReplicaInfo - Defines the characteristics of a replica.
// Modeled after SNIA Swordfish schema ReplicaInfo
type ReplicaInfo struct {
	// True if consistency is enabled.
	ConsistencyEnabled bool `json:"ConsistencyEnabled,omitempty"`

	// The current state of consistency.
	ConsistencyState ConsistencyState `json:"ConsistencyState,omitempty"`

	// The current status of consistency.
	ConsistencyStatus ConsistencyStatus `json:"ConsistencyStatus,omitempty"`

	// Indicates the consistency type used by the source and its associated target group.
	ConsistencyType ConsistencyType `json:"ConsistencyType,omitempty"`

	// A pointer to the DataProtection line of service element that describes this replica.
	DataProtectionLineOfService map[string]interface{} `json:"DataProtectionLineOfService,omitempty"`

	// If true, the storage array tells host to stop sending data to source element if copying to a remote element fails.
	FailedCopyStopsHostIO bool `json:"FailedCopyStopsHostIO,omitempty"`

	// Specifies the percent of the work completed to reach synchronization.
	PercentSynced int `json:"PercentSynced,omitempty"`

	// Deprecated - Use Source Replica. The resource that is the source of this replica.
	Replica map[string]interface{} `json:"Replica,omitempty"`

	// The priority of background copy engine I/O to be managed relative to host I/O operations during a sequential background copy operation.
	ReplicaPriority ReplicaPriority `json:"ReplicaPriority,omitempty"`

	// The status of the session with respect to Replication activity.
	ReplicaProgressStatus ReplicaProgressStatus `json:"ReplicaProgressStatus,omitempty"`

	// This property specifies whether the source, the target, or both elements are read only to the host.
	ReplicaReadOnlyAccess ReplicaReadOnlyAccess `json:"ReplicaReadOnlyAccess,omitempty"`

	// Describes whether the copy operation continues after a broken link is restored.
	ReplicaRecoveryMode ReplicaRecoveryMode `json:"ReplicaRecoveryMode,omitempty"`

	// The source or target role of this replica.
	ReplicaRole ReplicaRole `json:"ReplicaRole,omitempty"`

	// Applies to Adaptive mode and it describes maximum number of bytes the SyncedElement (target) can be out of sync.
	ReplicaSkewBytes int `json:"ReplicaSkewBytes,omitempty"`

	// ReplicaState describes the state of the relationship with respect to Replication activity.
	ReplicaState ReplicaState `json:"ReplicaState,omitempty"`

	// ReplicaType describes the intended outcome of the replication.
	ReplicaType map[string]interface{} `json:"ReplicaType,omitempty"`

	// Describes whether the target elements will be updated synchronously or asynchronously.
	ReplicaUpdateMode map[string]interface{} `json:"ReplicaUpdateMode,omitempty"`

	// The last requested or desired state for the relationship.
	RequestedReplicaState ReplicaState `json:"RequestedReplicaState,omitempty"`

	// Synchronization is maintained.
	SyncMaintained bool `json:"SyncMaintained,omitempty"`

	// This property specifies whether the source, the target, or both elements involved in a copy operation are undiscovered.
	UndiscoveredElement UndiscoveredElement `json:"UndiscoveredElement,omitempty"`

	// Specifies when point-in-time copy was taken or when the replication relationship is activated, reactivated, resumed or re-established.
	WhenActivated string `json:"WhenActivated,omitempty"`

	// Specifies when the replication relationship is deactivated.
	WhenDeactivated string `json:"WhenDeactivated,omitempty"`

	// Specifies when the replication relationship is established.
	WhenEstablished string `json:"WhenEstablished,omitempty"`

	// Specifies when the replication relationship is suspended.
	WhenSuspended string `json:"WhenSuspended,omitempty"`

	// The point in time that the Elements were synchronized.
	WhenSynced string `json:"WhenSynced,omitempty"`

	// Specifies when the replication relationship is synchronized.
	WhenSynchronized string `json:"WhenSynchronized,omitempty"`

}
