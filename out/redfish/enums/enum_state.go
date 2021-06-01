/* -----------------------------------------------------------------
* enum_state.go -
*
* DMTF Redfish State enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type State string

const (
	// This function or resource is enabled.
	State_ENABLED State = "Enabled"

	// This function or resource is disabled.
	State_DISABLED State = "Disabled"

	// This function or resource is enabled but awaits an external action to activate it.
	State_STANDBY_OFFLINE State = "StandbyOffline"

	// This function or resource is part of a redundancy set and awaits a failover or other external action to activate it.
	State_STANDBY_SPARE State = "StandbySpare"

	// This function or resource is undergoing testing, or is in the process of capturing information for debugging.
	State_IN_TEST State = "InTest"

	// This function or resource is starting.
	State_STARTING State = "Starting"

	// This function or resource is either not present or detected.
	State_ABSENT State = "Absent"

	// This function or resource is present but cannot be used.
	State_UNAVAILABLE_OFFLINE State = "UnavailableOffline"

	// The element does not process any commands but queues new requests.
	State_DEFERRING State = "Deferring"

	// The element is enabled but only processes a restricted set of commands.
	State_QUIESCED State = "Quiesced"

	// The element is updating and might be unavailable or degraded.
	State_UPDATING State = "Updating"

	// The element quality is within the acceptable range of operation.
	State_QUALIFIED State = "Qualified"

)
