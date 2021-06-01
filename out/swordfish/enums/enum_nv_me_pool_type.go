/* -----------------------------------------------------------------
* enum_nv_me_pool_type.go -
*
* SNIA Swordfish NVMePoolType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type NVMePoolType string

const (
	// This pool is of type EnduranceGroup, used by NVMe devices.
	NVMePoolType_ENDURANCE_GROUP NVMePoolType = "EnduranceGroup"

	// This pool is of type NVMSet, used by NVMe devices.
	NVMePoolType_NVM_SET NVMePoolType = "NVMSet"

)
