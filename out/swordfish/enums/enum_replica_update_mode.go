/* -----------------------------------------------------------------
* enum_replica_update_mode.go -
*
* SNIA Swordfish ReplicaUpdateMode enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaUpdateMode string

const (
	// Active-Active (i.e. bidirectional) synchronous updates.
	ReplicaUpdateMode_ACTIVE ReplicaUpdateMode = "Active"

	// Synchronous updates.
	ReplicaUpdateMode_SYNCHRONOUS ReplicaUpdateMode = "Synchronous"

	// Asynchronous updates.
	ReplicaUpdateMode_ASYNCHRONOUS ReplicaUpdateMode = "Asynchronous"

	// Allows implementation to switch between synchronous and asynchronous modes.
	ReplicaUpdateMode_ADAPTIVE ReplicaUpdateMode = "Adaptive"

)
