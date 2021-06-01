/* -----------------------------------------------------------------
* manager_service.go -
*
* DMTF Redfish ManagerService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ManagerService - Used for describing services like Serial Console, Command Shell or Graphical Console.
// Modeled after DMTF Redfish schema ManagerService
type ManagerService struct {
	// Indicates the maximum number of service sessions, regardless of protocol, this manager is able to support.
	MaxConcurrentSessions int `json:"MaxConcurrentSessions,omitempty"`

	// Indicates if the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

}
