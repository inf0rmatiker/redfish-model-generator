/* -----------------------------------------------------------------
 * enum_power_equipment_type.go -
 *
 * DMTF Redfish PowerEquipmentType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type PowerEquipmentType string

const (
	// A power distribution unit providing outlets for a rack or similar quantity of devices.
	PowerEquipmentType_RACK_PDU PowerEquipmentType = "RackPDU"

	// A power distribution unit providing feeder circuits for further power distribution.
	PowerEquipmentType_FLOOR_PDU PowerEquipmentType = "FloorPDU"

	// A manual power transfer switch.
	PowerEquipmentType_MANUAL_TRANSFER_SWITCH PowerEquipmentType = "ManualTransferSwitch"

	// An automatic power transfer switch.
	PowerEquipmentType_AUTOMATIC_TRANSFER_SWITCH PowerEquipmentType = "AutomaticTransferSwitch"

	// Electrical switchgear.
	PowerEquipmentType_SWITCHGEAR PowerEquipmentType = "Switchgear"

)
