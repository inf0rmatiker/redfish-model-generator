/* -----------------------------------------------------------------
* enum_physical_context.go -
*
* DMTF Redfish PhysicalContext enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PhysicalContext string

const (
	// The room
	PhysicalContext_ROOM PhysicalContext = "Room"

	// The intake point of the chassis
	PhysicalContext_INTAKE PhysicalContext = "Intake"

	// The exhaust point of the chassis
	PhysicalContext_EXHAUST PhysicalContext = "Exhaust"

	// The front of the chassis
	PhysicalContext_FRONT PhysicalContext = "Front"

	// The back of the chassis
	PhysicalContext_BACK PhysicalContext = "Back"

	// The upper portion of the chassis
	PhysicalContext_UPPER PhysicalContext = "Upper"

	// The lower portion of the chassis
	PhysicalContext_LOWER PhysicalContext = "Lower"

	// A Processor (CPU)
	PhysicalContext_CPU PhysicalContext = "CPU"

	// A Graphics Processor (GPU)
	PhysicalContext_GPU PhysicalContext = "GPU"

	// A backplane within the chassis
	PhysicalContext_BACKPLANE PhysicalContext = "Backplane"

	// The system board (PCB)
	PhysicalContext_SYSTEM_BOARD PhysicalContext = "SystemBoard"

	// A power supply
	PhysicalContext_POWER_SUPPLY PhysicalContext = "PowerSupply"

	// A voltage regulator device
	PhysicalContext_VOLTAGE_REGULATOR PhysicalContext = "VoltageRegulator"

	// A storage device
	PhysicalContext_STORAGE_DEVICE PhysicalContext = "StorageDevice"

	// A networking device
	PhysicalContext_NETWORKING_DEVICE PhysicalContext = "NetworkingDevice"

	// Within a compute bay
	PhysicalContext_COMPUTE_BAY PhysicalContext = "ComputeBay"

	// Within a storage bay
	PhysicalContext_STORAGE_BAY PhysicalContext = "StorageBay"

	// Within a networking bay
	PhysicalContext_NETWORK_BAY PhysicalContext = "NetworkBay"

	// Within an expansion bay
	PhysicalContext_EXPANSION_BAY PhysicalContext = "ExpansionBay"

	// Within a power supply bay
	PhysicalContext_POWER_SUPPLY_BAY PhysicalContext = "PowerSupplyBay"

)
