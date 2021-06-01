/* -----------------------------------------------------------------
* enum_failure_domain_scope.go -
*
* SNIA Swordfish FailureDomainScope enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type FailureDomainScope string

const (
	// A CPU/memory complex.
	FailureDomainScope_SERVER FailureDomainScope = "Server"

	// A container for Servers, Networking, and Storage.
	FailureDomainScope_RACK FailureDomainScope = "Rack"

	// A set of Racks that share common infrastructure.
	FailureDomainScope_RACK_GROUP FailureDomainScope = "RackGroup"

	// An adjacent set of racks.
	FailureDomainScope_ROW FailureDomainScope = "Row"

	// A co-located set of servers, including network and storage that share communication, power, or cooling infrastructure.
	FailureDomainScope_DATACENTER FailureDomainScope = "Datacenter"

	// A geographical or politically isolated set of resources.
	FailureDomainScope_REGION FailureDomainScope = "Region"

)
