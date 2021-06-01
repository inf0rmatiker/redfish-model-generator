/* -----------------------------------------------------------------
* enum_replica_fault_domain.go -
*
* SNIA Swordfish ReplicaFaultDomain enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ReplicaFaultDomain string

const (
	// Local indicates that the source and target replicas are contained within a single fault domain.
	ReplicaFaultDomain_LOCAL ReplicaFaultDomain = "Local"

	// Remote indicates that the source and target replicas are in separate fault domains.
	ReplicaFaultDomain_REMOTE ReplicaFaultDomain = "Remote"

)
