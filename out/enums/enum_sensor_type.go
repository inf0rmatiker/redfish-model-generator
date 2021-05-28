/* -----------------------------------------------------------------
 * enum_sensor_type.go -
 *
 * DMTF Redfish SensorType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SensorType string

const (
	// A platform security sensor.
	SensorType_PLATFORM _SECURITY _VIOLATION _ATTEMPT SensorType = "Platform Security Violation Attempt"

	// A temperature sensor.
	SensorType_TEMPERATURE SensorType = "Temperature"

	// A voltage sensor.
	SensorType_VOLTAGE SensorType = "Voltage"

	// A current sensor.
	SensorType_CURRENT SensorType = "Current"

	// A fan sensor.
	SensorType_FAN SensorType = "Fan"

	// A physical security sensor.
	SensorType_PHYSICAL _CHASSIS _SECURITY SensorType = "Physical Chassis Security"

	// A sensor for a processor.
	SensorType_PROCESSOR SensorType = "Processor"

	// A sensor for a power supply or DC-to-DC converter.
	SensorType_POWER _SUPPLY / _CONVERTER SensorType = "Power Supply / Converter"

	// A sensor for a power unit.
	SensorType_POWER_UNIT SensorType = "PowerUnit"

	// A sensor for a cooling device.
	SensorType_COOLING_DEVICE SensorType = "CoolingDevice"

	// A sensor for a miscellaneous analog sensor.
	SensorType_OTHER _UNITS-BASED _SENSOR SensorType = "Other Units-based Sensor"

	// A sensor for a memory device.
	SensorType_MEMORY SensorType = "Memory"

	// A sensor for a drive slot or bay.
	SensorType_DRIVE _SLOT/_BAY SensorType = "Drive Slot/Bay"

	// A sensor for a POST memory resize event.
	SensorType_POST _MEMORY _RESIZE SensorType = "POST Memory Resize"

	// A sensor for a system firmware progress event.
	SensorType_SYSTEM _FIRMWARE _PROGRESS SensorType = "System Firmware Progress"

	// A sensor for the event log.
	SensorType_EVENT _LOGGING _DISABLED SensorType = "Event Logging Disabled"

	// A sensor for a system event.
	SensorType_SYSTEM _EVENT SensorType = "System Event"

	// A sensor for a critical interrupt event.
	SensorType_CRITICAL _INTERRUPT SensorType = "Critical Interrupt"

	// A sensor for a button or switch.
	SensorType_BUTTON/_SWITCH SensorType = "Button/Switch"

	// A sensor for a module or board.
	SensorType_MODULE/_BOARD SensorType = "Module/Board"

	// A sensor for a microcontroller or coprocessor.
	SensorType_MICROCONTROLLER/_COPROCESSOR SensorType = "Microcontroller/Coprocessor"

	// A sensor for an add-in card.
	SensorType_ADD-IN _CARD SensorType = "Add-in Card"

	// A sensor for a chassis.
	SensorType_CHASSIS SensorType = "Chassis"

	// A sensor for a chipset.
	SensorType_CHIP_SET SensorType = "ChipSet"

	// A sensor for another type of FRU.
	SensorType_OTHER FRU SensorType = "Other FRU"

	// A sensor for a cable or interconnect device type.
	SensorType_CABLE/_INTERCONNECT SensorType = "Cable/Interconnect"

	// A sensor for a terminator.
	SensorType_TERMINATOR SensorType = "Terminator"

	// A sensor for a system boot or restart event.
	SensorType_SYSTEM_BOOT/_RESTART SensorType = "SystemBoot/Restart"

	// A sensor for a boot error event.
	SensorType_BOOT _ERROR SensorType = "Boot Error"

	// A sensor for a base OS boot or installation status event.
	SensorType_BASE_OS_BOOT/_INSTALLATION_STATUS SensorType = "BaseOSBoot/InstallationStatus"

	// A sensor for an OS stop or shutdown event
	SensorType_OS _STOP/_SHUTDOWN SensorType = "OS Stop/Shutdown"

	// A sensor for a slot or connector.
	SensorType_SLOT/_CONNECTOR SensorType = "Slot/Connector"

	// A sensor for an ACPI power state event.
	SensorType_SYSTEM ACPI _POWER_STATE SensorType = "System ACPI PowerState"

	// A sensor for a watchdog event.
	SensorType_WATCHDOG SensorType = "Watchdog"

	// A sensor for a platform alert event.
	SensorType_PLATFORM _ALERT SensorType = "Platform Alert"

	// A sensor for an entity presence event.
	SensorType_ENTITY _PRESENCE SensorType = "Entity Presence"

	// A sensor for a monitor ASIC or IC.
	SensorType_MONITOR ASIC/IC SensorType = "Monitor ASIC/IC"

	// A sensor for a LAN device.
	SensorType_LAN SensorType = "LAN"

	// A sensor for a management subsystem health event.
	SensorType_MANAGEMENT _SUBSYSTEM _HEALTH SensorType = "Management Subsystem Health"

	// A sensor for a battery.
	SensorType_BATTERY SensorType = "Battery"

	// A sensor for a session audit event.
	SensorType_SESSION _AUDIT SensorType = "Session Audit"

	// A sensor for a version change event.
	SensorType_VERSION _CHANGE SensorType = "Version Change"

	// A sensor for a FRU state event.
	SensorType_FRU_STATE SensorType = "FRUState"

	// An OEM-defined sensor.
	SensorType_OEM SensorType = "OEM"

)
