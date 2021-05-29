/* -----------------------------------------------------------------
* watchdog_timer.go -
*
* DMTF Redfish WatchdogTimer resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes the host watchdog timer functionality for this system.
type WatchdogTimer struct {
	// An indication of whether a user has enabled the host watchdog timer functionality.  This property indicates only that a user has enabled the timer.  To activate the timer, installation of additional host-based software is necessary; an update to this property does not initiate the timer.
	FunctionEnabled bool `json:"FunctionEnabled"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The action to perform when the watchdog timer reaches its timeout value.
	TimeoutAction WatchdogTimeoutActions `json:"TimeoutAction"`

	// The action to perform when the watchdog timer is close to reaching its timeout value.  This action typically occurs from three to ten seconds before to the timeout value, but the exact timing is dependent on the implementation.
	WarningAction WatchdogWarningActions `json:"WarningAction,omitempty"`

}
