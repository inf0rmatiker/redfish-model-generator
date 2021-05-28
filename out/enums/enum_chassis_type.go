/* -----------------------------------------------------------------
 * enum_chassis_type.go -
 *
 * DMTF Redfish ChassisType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ChassisType string

const (
	// An equipment rack, typically a 19-inch wide freestanding unit.
	ChassisType_RACK ChassisType = "Rack"

	// An enclosed or semi-enclosed, typically vertically-oriented, system chassis that must be plugged into a multi-system chassis to function normally.
	ChassisType_BLADE ChassisType = "Blade"

	// A generic term for a chassis that does not fit any other description.
	ChassisType_ENCLOSURE ChassisType = "Enclosure"

	// A single, free-standing system, commonly called a tower or desktop chassis.
	ChassisType_STAND_ALONE ChassisType = "StandAlone"

	// A single-system chassis designed specifically for mounting in an equipment rack.
	ChassisType_RACK_MOUNT ChassisType = "RackMount"

	// A loose device or circuit board intended to be installed in a system or other enclosure.
	ChassisType_CARD ChassisType = "Card"

	// A small self-contained system intended to be plugged into a multi-system chassis.
	ChassisType_CARTRIDGE ChassisType = "Cartridge"

	// A collection of equipment racks.
	ChassisType_ROW ChassisType = "Row"

	// A collection of equipment racks in a large, likely transportable, container.
	ChassisType_POD ChassisType = "Pod"

	// A chassis that expands the capabilities or capacity of another chassis.
	ChassisType_EXPANSION ChassisType = "Expansion"

	// A chassis that mates mechanically with another chassis to expand its capabilities or capacity.
	ChassisType_SIDECAR ChassisType = "Sidecar"

	// A logical division or portion of a physical chassis that contains multiple devices or systems that cannot be physically separated.
	ChassisType_ZONE ChassisType = "Zone"

	// An enclosed or semi-enclosed, system chassis that must be plugged into a multi-system chassis to function normally similar to a blade type chassis.
	ChassisType_SLED ChassisType = "Sled"

	// An enclosed or semi-enclosed, typically horizontally-oriented, system chassis that must be plugged into a multi-system chassis to function normally.
	ChassisType_SHELF ChassisType = "Shelf"

	// An enclosed or semi-enclosed, typically horizontally-oriented, system chassis that can be slid into a multi-system chassis.
	ChassisType_DRAWER ChassisType = "Drawer"

	// A small, typically removable, chassis or card that contains devices for a particular subsystem or function.
	ChassisType_MODULE ChassisType = "Module"

	// A small chassis, card, or device that contains devices for a particular subsystem or function.
	ChassisType_COMPONENT ChassisType = "Component"

	// A chassis in a drive form factor with IP-based network connections.
	ChassisType_IP_BASED_DRIVE ChassisType = "IPBasedDrive"

	// A group of racks that form a single entity or share infrastructure.
	ChassisType_RACK_GROUP ChassisType = "RackGroup"

	// A chassis that encloses storage.
	ChassisType_STORAGE_ENCLOSURE ChassisType = "StorageEnclosure"

	// A chassis that does not fit any of these definitions.
	ChassisType_OTHER ChassisType = "Other"

)
