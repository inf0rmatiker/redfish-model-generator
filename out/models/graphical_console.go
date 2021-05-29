/* -----------------------------------------------------------------
* graphical_console.go -
*
* DMTF Redfish GraphicalConsole resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The information about a graphical console service that this manager provides.
type GraphicalConsole struct {
	// This property enumerates the graphical console connection types that the implementation allows.
	ConnectTypesSupported array `json:"ConnectTypesSupported,omitempty"`

	// The maximum number of service sessions, regardless of protocol, that this manager can support.
	MaxConcurrentSessions int `json:"MaxConcurrentSessions,omitempty"`

	// An indication of whether the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

}
