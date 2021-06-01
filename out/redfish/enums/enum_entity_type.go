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
	// The entity is a storage initator. The EntityLink property (if present) should be a Storage.StorageController entity.
	EntityType_STORAGE_INITIATOR EntityType = "StorageInitiator"

	// The entity is a PCI(e) root complex. The EntityLink property (if present) should be a ComputerSystem.ComputerSystem entity.
	EntityType_ROOT_COMPLEX EntityType = "RootComplex"

	// The entity is a network controller. The EntityLink property (if present) should be an EthernetInterface.EthernetInterface entity.
	EntityType_NETWORK_CONTROLLER EntityType = "NetworkController"

	// The entity is a disk drive. The EntityLink property (if present) should be a Drive.Drive entity.
	EntityType_DRIVE EntityType = "Drive"

	// The entity is a storage expander. The EntityLink property (if present) should be a Chassis.Chassis entity.
	EntityType_STORAGE_EXPANDER EntityType = "StorageExpander"

	// The entity is a display controller.
	EntityType_DISPLAY_CONTROLLER EntityType = "DisplayController"

	// The entity is a PCI(e) bridge.
	EntityType_BRIDGE EntityType = "Bridge"

	// The entity is a processor device.
	EntityType_PROCESSOR EntityType = "Processor"

	// The entity is a volume. The EntityLink property (if present) should be a Volume.Volume entity.
	EntityType_VOLUME EntityType = "Volume"

)
