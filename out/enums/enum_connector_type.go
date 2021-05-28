/* -----------------------------------------------------------------
 * enum_connector_type.go -
 *
 * DMTF Redfish ConnectorType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ConnectorType string

const (
	// An RJ45 connector.
	ConnectorType_RJ45 ConnectorType = "RJ45"

	// An RJ11 connector.
	ConnectorType_RJ11 ConnectorType = "RJ11"

	// A DB9 Female connector.
	ConnectorType_DB9 _FEMALE ConnectorType = "DB9 Female"

	// A DB9 Male connector.
	ConnectorType_DB9 _MALE ConnectorType = "DB9 Male"

	// A DB25 Female connector.
	ConnectorType_DB25 _FEMALE ConnectorType = "DB25 Female"

	// A DB25 Male connector.
	ConnectorType_DB25 _MALE ConnectorType = "DB25 Male"

	// A USB connector.
	ConnectorType_USB ConnectorType = "USB"

	// A mUSB connector.
	ConnectorType_M_USB ConnectorType = "mUSB"

	// A uUSB connector.
	ConnectorType_U_USB ConnectorType = "uUSB"

)
