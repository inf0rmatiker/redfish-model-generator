/* -----------------------------------------------------------------
* virtual_media.go -
*
* DMTF Redfish VirtualMedia resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VirtualMedia - This resource allows monitoring and control of an instance of virtual media (e.g. a remote CD, DVD, or USB device) functionality provided by a Manager for a system or device.
// Modeled after DMTF Redfish schema VirtualMedia
type VirtualMedia struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The current image name
	ImageName string `json:"ImageName,omitempty"`

	// A URI providing the location of the selected image
	Image string `json:"Image,omitempty"`

	// This is the media types supported as virtual media.
	MediaTypes []MediaType `json:"MediaTypes,omitempty"`

	// Current virtual media connection methods
	ConnectedVia ConnectedVia `json:"ConnectedVia,omitempty"`

	// Indicates if virtual media is inserted in the virtual device.
	Inserted bool `json:"Inserted,omitempty"`

	// Indicates the media is write protected.
	WriteProtected bool `json:"WriteProtected,omitempty"`

}
