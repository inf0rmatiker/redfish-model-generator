/* -----------------------------------------------------------------
* io_connectivity_line_of_service.go -
*
* SNIA Swordfish IOConnectivityLineOfService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IOConnectivityLineOfService - A service option within the IO Connectivity line of service.
// Modeled after SNIA Swordfish schema IOConnectivityLineOfService
type IOConnectivityLineOfService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// SupportedAccessProtocols.
	AccessProtocols []Protocol `json:"AccessProtocols,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The maximum Bandwidth in bytes per second that a connection can support.
	MaxBytesPerSecond int `json:"MaxBytesPerSecond,omitempty"`

	// The maximum supported IOs per second that the connection will support for the selected access protocol.
	MaxIOPS int `json:"MaxIOPS,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
