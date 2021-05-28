/* -----------------------------------------------------------------
 * enum_durable_name_format.go -
 *
 * DMTF Redfish DurableNameFormat enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DurableNameFormat string

const (
	// The Name Address Authority (NAA) format.
	DurableNameFormat_NAA DurableNameFormat = "NAA"

	// The iSCSI Qualified Name (iQN).
	DurableNameFormat_I_QN DurableNameFormat = "iQN"

	// The Fibre Channel (FC) World Wide Name (WWN).
	DurableNameFormat_FC_WWN DurableNameFormat = "FC_WWN"

	// The Universally Unique Identifier (UUID).
	DurableNameFormat_UUID DurableNameFormat = "UUID"

	// The IEEE-defined 64-bit Extended Unique Identifier (EUI).
	DurableNameFormat_EUI DurableNameFormat = "EUI"

	// The NVMe Qualified Name (NQN).
	DurableNameFormat_NQN DurableNameFormat = "NQN"

	// The NVM Namespace Identifier (NSID).
	DurableNameFormat_NSID DurableNameFormat = "NSID"

)
