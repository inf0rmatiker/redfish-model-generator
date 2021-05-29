/* -----------------------------------------------------------------
* command_shell.go -
*
* DMTF Redfish CommandShell resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The information about a command shell service that this manager provides.
type CommandShell struct {
	// This property enumerates the command shell connection types that the implementation allows.
	ConnectTypesSupported array `json:"ConnectTypesSupported,omitempty"`

	// The maximum number of service sessions, regardless of protocol, that this manager can support.
	MaxConcurrentSessions int `json:"MaxConcurrentSessions,omitempty"`

	// An indication of whether the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

}
