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
	// Name Address Authority Format.
	DurableNameFormat_NAA DurableNameFormat = "NAA"

	// iSCSI Qualified Name.
	DurableNameFormat_I_QN DurableNameFormat = "iQN"

	// Fibre Channel World Wide Name.
	DurableNameFormat_FC_WWN DurableNameFormat = "FC_WWN"

	// Universally Unique Identifier.
	DurableNameFormat_UUID DurableNameFormat = "UUID"

	// IEEE-defined 64-bit Extended Unique Identifier.
	DurableNameFormat_EUI DurableNameFormat = "EUI"

	// NVMe Qualified Name.
	DurableNameFormat_NQN DurableNameFormat = "NQN"

	// NVM Namespace Identifier.
	DurableNameFormat_NSID DurableNameFormat = "NSID"

)
