/* -----------------------------------------------------------------
* storage_controller_links.go -
*
* DMTF Redfish StorageControllerLinks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// StorageControllerLinks - Contains references to other resources that are related to this resource.
// Modeled after DMTF Redfish schema StorageControllerLinks
type StorageControllerLinks struct {
	// An array of references to the endpoints that connect to this controller.
	Endpoints []Endpoint `json:"Endpoints,omitempty"`

	EndpointsOdataCount int `json:"Endpoints@odata.count,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
