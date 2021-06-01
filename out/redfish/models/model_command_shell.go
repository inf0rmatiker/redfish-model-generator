/* -----------------------------------------------------------------
* command_shell.go -
*
* DMTF Redfish CommandShell resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CommandShell - Information about a Command Shell service provided by a manager.
// Modeled after DMTF Redfish schema CommandShell
type CommandShell struct {
	// This object is used to enumerate the Command Shell connection types allowed by the implementation.
	ConnectTypesSupported []CommandConnectTypesSupported `json:"ConnectTypesSupported,omitempty"`

	// Indicates the maximum number of service sessions, regardless of protocol, this manager is able to support.
	MaxConcurrentSessions int `json:"MaxConcurrentSessions,omitempty"`

	// Indicates if the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

}
