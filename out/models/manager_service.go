/* -----------------------------------------------------------------
* manager_service.go -
*
* DMTF Redfish ManagerService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The manager services, such as serial console, command shell, or graphical console service.
type ManagerService struct {
	// The maximum number of service sessions, regardless of protocol, that this manager can support.
	MaxConcurrentSessions int `json:"MaxConcurrentSessions,omitempty"`

	// An indication of whether the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

}
