/* -----------------------------------------------------------------
* enum_replica_type.go -
*
* SNIA Swordfish ReplicaType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaType string

const (
	// Create and maintain a copy of the source.
	ReplicaType_MIRROR ReplicaType = "Mirror"

	// Create a point in time, virtual copy of the source.
	ReplicaType_SNAPSHOT ReplicaType = "Snapshot"

	// Create a point in time, full copy the source.
	ReplicaType_CLONE ReplicaType = "Clone"

	// Create a token based clone.
	ReplicaType_TOKENIZED_CLONE ReplicaType = "TokenizedClone"

)
