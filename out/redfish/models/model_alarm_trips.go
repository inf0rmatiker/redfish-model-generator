/* -----------------------------------------------------------------
* alarm_trips.go -
*
* DMTF Redfish AlarmTrips resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// AlarmTrips - The alarm trip information about the memory.  These alarms are reset when the system resets.  Note that if they are re-discovered they can be reasserted.
// Modeled after DMTF Redfish schema AlarmTrips
type AlarmTrips struct {
	// An indication of whether an address parity error was detected that a retry could not correct.
	AddressParityError bool `json:"AddressParityError,omitempty"`

	// An indication of whether the correctable error threshold crossing alarm trip was detected.
	CorrectableECCError bool `json:"CorrectableECCError,omitempty"`

	// An indication of whether the spare block capacity crossing alarm trip was detected.
	SpareBlock bool `json:"SpareBlock,omitempty"`

	// An indication of whether a temperature threshold alarm trip was detected.
	Temperature bool `json:"Temperature,omitempty"`

	// An indication of whether the uncorrectable error threshold alarm trip was detected.
	UncorrectableECCError bool `json:"UncorrectableECCError,omitempty"`

}
