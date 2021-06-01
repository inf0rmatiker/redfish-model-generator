/* -----------------------------------------------------------------
* hosted_services.go -
*
* DMTF Redfish HostedServices resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// HostedServices - The services that might be running or installed on the system.
// Modeled after DMTF Redfish schema HostedServices
type HostedServices struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to a collection of storage services that this computer system supports.
	StorageServices map[string]interface{} `json:"StorageServices,omitempty"`

}
