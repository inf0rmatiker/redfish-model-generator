/* -----------------------------------------------------------------
* graphical_console.go -
*
* DMTF Redfish GraphicalConsole resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// GraphicalConsole - Information about a Graphical Console service provided by a manager.
// Modeled after DMTF Redfish schema GraphicalConsole
type GraphicalConsole struct {
	// This object is used to enumerate the Graphical Console connection types allowed by the implementation.
	ConnectTypesSupported []GraphicalConnectTypesSupported `json:"ConnectTypesSupported,omitempty"`

	// Indicates the maximum number of service sessions, regardless of protocol, this manager is able to support.
	MaxConcurrentSessions int `json:"MaxConcurrentSessions,omitempty"`

	// Indicates if the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

}
