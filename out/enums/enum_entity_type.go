/* -----------------------------------------------------------------
 * enum_entity_type.go -
 *
 * DMTF Redfish EntityType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type EntityType string

const (
	// The entity is a storage initiator.
	EntityType_STORAGE_INITIATOR EntityType = "StorageInitiator"

	// The entity is a PCI(e) root complex.
	EntityType_ROOT_COMPLEX EntityType = "RootComplex"

	// The entity is a network controller.
	EntityType_NETWORK_CONTROLLER EntityType = "NetworkController"

	// The entity is a drive.
	EntityType_DRIVE EntityType = "Drive"

	// The entity is a storage expander.
	EntityType_STORAGE_EXPANDER EntityType = "StorageExpander"

	// The entity is a display controller.
	EntityType_DISPLAY_CONTROLLER EntityType = "DisplayController"

	// The entity is a PCI(e) bridge.
	EntityType_BRIDGE EntityType = "Bridge"

	// The entity is a processor.
	EntityType_PROCESSOR EntityType = "Processor"

	// The entity is a volume.
	EntityType_VOLUME EntityType = "Volume"

	// The entity is an acceleration function realized through a device, such as an FPGA.
	EntityType_ACCELERATION_FUNCTION EntityType = "AccelerationFunction"

	// The entity is a media controller.
	EntityType_MEDIA_CONTROLLER EntityType = "MediaController"

	// The entity is a memory chunk.
	EntityType_MEMORY_CHUNK EntityType = "MemoryChunk"

	// The entity is a switch, not an expander.  Use `Expander` for expanders.
	EntityType_SWITCH EntityType = "Switch"

	// The entity is a fabric bridge.
	EntityType_FABRIC_BRIDGE EntityType = "FabricBridge"

	// The entity is a manager.
	EntityType_MANAGER EntityType = "Manager"

	// The entity is a storage subsystem.
	EntityType_STORAGE_SUBSYSTEM EntityType = "StorageSubsystem"

)
