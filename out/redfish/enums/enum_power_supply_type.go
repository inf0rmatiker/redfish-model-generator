/* -----------------------------------------------------------------
* enum_power_supply_type.go -
*
* DMTF Redfish PowerSupplyType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PowerSupplyType string

const (
	// The power supply type cannot be determined.
	PowerSupplyType_UNKNOWN PowerSupplyType = "Unknown"

	// Alternating Current (AC) power supply.
	PowerSupplyType_AC PowerSupplyType = "AC"

	// Direct Current (DC) power supply.
	PowerSupplyType_DC PowerSupplyType = "DC"

	// Power Supply supports both DC or AC.
	PowerSupplyType_A_COR_DC PowerSupplyType = "ACorDC"

)
