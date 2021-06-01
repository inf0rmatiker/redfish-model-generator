/* -----------------------------------------------------------------
* enum_replica_priority.go -
*
* SNIA Swordfish ReplicaPriority enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaPriority string

const (
	// Copy engine I/O lower priority than host I/O.
	ReplicaPriority_LOW ReplicaPriority = "Low"

	// Copy engine I/O has the same priority as host I/O.
	ReplicaPriority_SAME ReplicaPriority = "Same"

	// Copy engine I/O has higher priority than host I/O.
	ReplicaPriority_HIGH ReplicaPriority = "High"

	// Copy operation to be performed as soon as possible, regardless of the host I/O requests.
	ReplicaPriority_URGENT ReplicaPriority = "Urgent"

)
