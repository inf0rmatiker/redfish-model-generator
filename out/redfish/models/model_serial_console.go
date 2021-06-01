/* -----------------------------------------------------------------
* serial_console.go -
*
* DMTF Redfish SerialConsole resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SerialConsole - Information about a Serial Console service provided by a manager.
// Modeled after DMTF Redfish schema SerialConsole
type SerialConsole struct {
	// This object is used to enumerate the Serial Console connection types allowed by the implementation.
	ConnectTypesSupported []SerialConnectTypesSupported `json:"ConnectTypesSupported,omitempty"`

	// Indicates the maximum number of service sessions, regardless of protocol, this manager is able to support.
	MaxConcurrentSessions int `json:"MaxConcurrentSessions,omitempty"`

	// Indicates if the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

}
