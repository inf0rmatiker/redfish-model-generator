/* -----------------------------------------------------------------
 * enum_slot_types.go -
 *
 * DMTF Redfish SlotTypes enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SlotTypes string

const (
	// Full-Length PCIe slot.
	SlotTypes_FULL_LENGTH SlotTypes = "FullLength"

	// Half-Length PCIe slot.
	SlotTypes_HALF_LENGTH SlotTypes = "HalfLength"

	// Low-Profile or Slim PCIe slot.
	SlotTypes_LOW_PROFILE SlotTypes = "LowProfile"

	// Mini PCIe slot.
	SlotTypes_MINI SlotTypes = "Mini"

	// PCIe M.2 slot.
	SlotTypes_M2 SlotTypes = "M2"

	// An OEM-specific slot.
	SlotTypes_OEM SlotTypes = "OEM"

	// Open Compute Project 3.0 small form factor slot.
	SlotTypes_OCP3_SMALL SlotTypes = "OCP3Small"

	// Open Compute Project 3.0 large form factor slot.
	SlotTypes_OCP3_LARGE SlotTypes = "OCP3Large"

	// U.2 / SFF-8639 slot or bay.
	SlotTypes_U2 SlotTypes = "U2"

)
