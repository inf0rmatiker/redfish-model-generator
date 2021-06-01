/* -----------------------------------------------------------------
* enum_write_hole_protection_policy_type.go -
*
* DMTF Redfish WriteHoleProtectionPolicyType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type WriteHoleProtectionPolicyType string

const (
	// The volume is not using any policy to address the write hole issue.
	WriteHoleProtectionPolicyType_OFF WriteHoleProtectionPolicyType = "Off"

	// The policy that uses separate block device for write-ahead logging to address write hole issue.
	WriteHoleProtectionPolicyType_JOURNALING WriteHoleProtectionPolicyType = "Journaling"

	// The policy that distributes additional log among the volume's capacity sources to address write hole issue.
	WriteHoleProtectionPolicyType_DISTRIBUTED_LOG WriteHoleProtectionPolicyType = "DistributedLog"

	// The policy that is Oem specific.
	WriteHoleProtectionPolicyType_OEM WriteHoleProtectionPolicyType = "Oem"

)
