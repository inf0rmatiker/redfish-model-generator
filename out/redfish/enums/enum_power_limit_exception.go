/* -----------------------------------------------------------------
* enum_power_limit_exception.go -
*
* DMTF Redfish PowerLimitException enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PowerLimitException string

const (
	// Take no action when the limit is exceeded.
	PowerLimitException_NO_ACTION PowerLimitException = "NoAction"

	// Turn the power off immediately when the limit is exceeded.
	PowerLimitException_HARD_POWER_OFF PowerLimitException = "HardPowerOff"

	// Log an event when the limit is exceeded, but take no further action.
	PowerLimitException_LOG_EVENT_ONLY PowerLimitException = "LogEventOnly"

	// Take an OEM-defined action.
	PowerLimitException_OEM PowerLimitException = "Oem"

)
