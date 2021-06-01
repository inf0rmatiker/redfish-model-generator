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
	ConnectorType_RJ45 ConnectorType = "RJ45"
	ConnectorType_RJ11 ConnectorType = "RJ11"
	ConnectorType_DB9 _FEMALE ConnectorType = "DB9 Female"
	ConnectorType_DB9 _MALE ConnectorType = "DB9 Male"
	ConnectorType_DB25 _FEMALE ConnectorType = "DB25 Female"
	ConnectorType_DB25 _MALE ConnectorType = "DB25 Male"
	ConnectorType_USB ConnectorType = "USB"
	ConnectorType_M_USB ConnectorType = "mUSB"
	ConnectorType_U_USB ConnectorType = "uUSB"
)
