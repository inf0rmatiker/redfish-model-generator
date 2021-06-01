/* -----------------------------------------------------------------
* enum_recovery_access_scope.go -
*
* SNIA Swordfish RecoveryAccessScope enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type RecoveryAccessScope string

const (
	// Active access to synchronous replicas.
	RecoveryAccessScope_ONLINE_ACTIVE RecoveryAccessScope = "OnlineActive"

	// Passive access to replicas via the same front-end interconnect.
	RecoveryAccessScope_ONLINE_PASSIVE RecoveryAccessScope = "OnlinePassive"

	// Access to replica via a different front-end interconnect. A restore step is required before recovery can commence.
	RecoveryAccessScope_NEARLINE RecoveryAccessScope = "Nearline"

	// No direct connection to the replica. (i.e. To a bunker containing backup media).
	RecoveryAccessScope_OFFLINE RecoveryAccessScope = "Offline"

)
