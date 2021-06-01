/* -----------------------------------------------------------------
* enum_device_class.go -
*
* DMTF Redfish DeviceClass enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type DeviceClass string

const (
	// An unclassified device.
	DeviceClass_UNCLASSIFIED_DEVICE DeviceClass = "UnclassifiedDevice"

	// A mass storage controller.
	DeviceClass_MASS_STORAGE_CONTROLLER DeviceClass = "MassStorageController"

	// A network controller.
	DeviceClass_NETWORK_CONTROLLER DeviceClass = "NetworkController"

	// A display controller.
	DeviceClass_DISPLAY_CONTROLLER DeviceClass = "DisplayController"

	// A multimedia controller.
	DeviceClass_MULTIMEDIA_CONTROLLER DeviceClass = "MultimediaController"

	// A memory controller.
	DeviceClass_MEMORY_CONTROLLER DeviceClass = "MemoryController"

	// A bridge.
	DeviceClass_BRIDGE DeviceClass = "Bridge"

	// A communication controller.
	DeviceClass_COMMUNICATION_CONTROLLER DeviceClass = "CommunicationController"

	// A generic system peripheral.
	DeviceClass_GENERIC_SYSTEM_PERIPHERAL DeviceClass = "GenericSystemPeripheral"

	// An input device controller.
	DeviceClass_INPUT_DEVICE_CONTROLLER DeviceClass = "InputDeviceController"

	// A docking station.
	DeviceClass_DOCKING_STATION DeviceClass = "DockingStation"

	// A processor.
	DeviceClass_PROCESSOR DeviceClass = "Processor"

	// A serial bus controller.
	DeviceClass_SERIAL_BUS_CONTROLLER DeviceClass = "SerialBusController"

	// A wireless controller.
	DeviceClass_WIRELESS_CONTROLLER DeviceClass = "WirelessController"

	// An intelligent controller.
	DeviceClass_INTELLIGENT_CONTROLLER DeviceClass = "IntelligentController"

	// A satellite communications controller.
	DeviceClass_SATELLITE_COMMUNICATIONS_CONTROLLER DeviceClass = "SatelliteCommunicationsController"

	// An encryption controller.
	DeviceClass_ENCRYPTION_CONTROLLER DeviceClass = "EncryptionController"

	// A signal processing controller.
	DeviceClass_SIGNAL_PROCESSING_CONTROLLER DeviceClass = "SignalProcessingController"

	// A processing accelerators.
	DeviceClass_PROCESSING_ACCELERATORS DeviceClass = "ProcessingAccelerators"

	// A non-essential instrumentation.
	DeviceClass_NON_ESSENTIAL_INSTRUMENTATION DeviceClass = "NonEssentialInstrumentation"

	// A coprocessor.
	DeviceClass_COPROCESSOR DeviceClass = "Coprocessor"

	// An unassigned class.
	DeviceClass_UNASSIGNED_CLASS DeviceClass = "UnassignedClass"

	// A other class.  The function Device Class Id needs to be verified.
	DeviceClass_OTHER DeviceClass = "Other"

)
