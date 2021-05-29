/* -----------------------------------------------------------------
* update_parameters.go -
*
* DMTF Redfish UpdateParameters resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The update parameters used with MultipartHttpPushUri software update.
type UpdateParameters struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of URIs that indicate where to apply the update image.
	Targets []string `json:"Targets,omitempty"`

}
