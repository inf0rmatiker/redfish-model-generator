/* -----------------------------------------------------------------
* enum_replica_role.go -
*
* SNIA Swordfish ReplicaRole enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaRole string

const (
	// The source element.
	ReplicaRole_SOURCE ReplicaRole = "Source"

	// The target element.
	ReplicaRole_TARGET ReplicaRole = "Target"

)
