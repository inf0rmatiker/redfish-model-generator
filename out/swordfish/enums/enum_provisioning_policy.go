/* -----------------------------------------------------------------
* enum_provisioning_policy.go -
*
* SNIA Swordfish ProvisioningPolicy enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ProvisioningPolicy string

const (
	// Storage is fully allocated.
	ProvisioningPolicy_FIXED ProvisioningPolicy = "Fixed"

	// Storage may be over allocated.
	ProvisioningPolicy_THIN ProvisioningPolicy = "Thin"

)
