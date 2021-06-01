/* -----------------------------------------------------------------
* enum_replica_read_only_access.go -
*
* SNIA Swordfish ReplicaReadOnlyAccess enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaReadOnlyAccess string

const (
	// The source element.
	ReplicaReadOnlyAccess_SOURCE_ELEMENT ReplicaReadOnlyAccess = "SourceElement"

	// The replica element.
	ReplicaReadOnlyAccess_REPLICA_ELEMENT ReplicaReadOnlyAccess = "ReplicaElement"

	// Both the source and the target elements are read only to the host.
	ReplicaReadOnlyAccess_BOTH ReplicaReadOnlyAccess = "Both"

)
