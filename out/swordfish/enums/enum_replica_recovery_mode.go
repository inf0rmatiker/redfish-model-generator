/* -----------------------------------------------------------------
* enum_replica_recovery_mode.go -
*
* SNIA Swordfish ReplicaRecoveryMode enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaRecoveryMode string

const (
	// Copy operation resumes automatically.
	ReplicaRecoveryMode_AUTOMATIC ReplicaRecoveryMode = "Automatic"

	// ReplicaState is set to Suspended after the link is restored. It is required to issue the Resume operation to continue.
	ReplicaRecoveryMode_MANUAL ReplicaRecoveryMode = "Manual"

)
