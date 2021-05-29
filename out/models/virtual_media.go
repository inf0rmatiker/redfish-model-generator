/* -----------------------------------------------------------------
* virtual_media.go -
*
* DMTF Redfish VirtualMedia resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The VirtualMedia schema contains properties related to the monitor and control of an instance of virtual media, such as a remote CD, DVD, or USB device.  A manager for a system or device provides virtual media functionality.
type VirtualMedia struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to a collection of server certificates for the server referenced by the Image property.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// The current virtual media connection method.
	ConnectedVia ConnectedVia `json:"ConnectedVia,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The URI of the location of the selected image.
	Image string `json:"Image,omitempty"`

	// The current image name.
	ImageName string `json:"ImageName,omitempty"`

	// An indication of whether virtual media is inserted into the virtual device.
	Inserted bool `json:"Inserted,omitempty"`

	// The media types supported as virtual media.
	MediaTypes array `json:"MediaTypes,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The password to access the Image parameter-specified URI.  This property is null in responses.
	Password string `json:"Password,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The transfer method to use with the Image.
	TransferMethod TransferMethod `json:"TransferMethod,omitempty"`

	// The network protocol to use with the image.
	TransferProtocolType TransferProtocolType `json:"TransferProtocolType,omitempty"`

	// The user name to access the Image parameter-specified URI.
	UserName string `json:"UserName,omitempty"`

	// An indication of whether the service will verify the certificate of the server referenced by the Image property prior to completing the remote media connection.
	VerifyCertificate bool `json:"VerifyCertificate,omitempty"`

	// An indication of whether the media is write-protected.
	WriteProtected bool `json:"WriteProtected,omitempty"`

}
