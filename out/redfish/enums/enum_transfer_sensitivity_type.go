/* -----------------------------------------------------------------
* enum_transfer_sensitivity_type.go -
*
* DMTF Redfish TransferSensitivityType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type TransferSensitivityType string

const (
	// High sensitivity for initiating a transfer.
	TransferSensitivityType_HIGH TransferSensitivityType = "High"

	// Medium sensitivity for initiating a transfer.
	TransferSensitivityType_MEDIUM TransferSensitivityType = "Medium"

	// Low sensitivity for initiating a transfer.
	TransferSensitivityType_LOW TransferSensitivityType = "Low"

)
