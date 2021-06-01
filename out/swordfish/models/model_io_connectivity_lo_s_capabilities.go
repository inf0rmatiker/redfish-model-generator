/* -----------------------------------------------------------------
* io_connectivity_lo_s_capabilities.go -
*
* SNIA Swordfish IOConnectivityLoSCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IOConnectivityLoSCapabilities - Describe IO Connectivity capabilities.
// Modeled after SNIA Swordfish schema IOConnectivityLoSCapabilities
type IOConnectivityLoSCapabilities struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// The maximum Bandwidth in bytes per second that a connection can support.
	MaxSupportedBytesPerSecond int `json:"MaxSupportedBytesPerSecond,omitempty"`

	// The maximum IOPS that a connection can support.
	MaxSupportedIOPS int `json:"MaxSupportedIOPS,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// SupportedAccessProtocols.
	SupportedAccessProtocols []Protocol `json:"SupportedAccessProtocols,omitempty"`

	// Collection of known and supported IOConnectivityLinesOfService.
	SupportedLinesOfService []IOConnectivityLineOfService `json:"SupportedLinesOfService,omitempty"`

	SupportedLinesOfServiceOdataCount int `json:"SupportedLinesOfService@odata.count,omitempty"`

}
