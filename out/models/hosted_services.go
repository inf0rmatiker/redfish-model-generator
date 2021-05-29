/* -----------------------------------------------------------------
* hosted_services.go -
*
* DMTF Redfish HostedServices resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The services that might be running or installed on the system.
type HostedServices struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to a collection of storage services that this computer system supports.
	StorageServices HostedStorageServices `json:"StorageServices,omitempty"`

}
