/* -----------------------------------------------------------------
* update_service.go -
*
* DMTF Redfish UpdateService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The UpdateService schema describes the update service and the properties for the service itself with links to collections of firmware and software inventory.  The update service also provides methods for updating software and firmware of the resources in a Redfish service.
type UpdateService struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An inventory of firmware.
	FirmwareInventory SoftwareInventoryCollection `json:"FirmwareInventory,omitempty"`

	// The URI used to perform an HTTP or HTTPS push update to the update service.  The format of the message is vendor-specific.
	HttpPushUri string `json:"HttpPushUri,omitempty"`

	// The options for HttpPushUri-provided software updates.
	HttpPushUriOptions HttpPushUriOptions `json:"HttpPushUriOptions,omitempty"`

	// An indication of whether a client has reserved the HttpPushUriOptions properties for software updates.
	HttpPushUriOptionsBusy bool `json:"HttpPushUriOptionsBusy,omitempty"`

	// An array of URIs that indicate where to apply the update image.
	HttpPushUriTargets []string `json:"HttpPushUriTargets,omitempty"`

	// An indication of whether any client has reserved the HttpPushUriTargets property.
	HttpPushUriTargetsBusy bool `json:"HttpPushUriTargetsBusy,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The maximum size in bytes of the software update image that this service supports.
	MaxImageSizeBytes int `json:"MaxImageSizeBytes,omitempty"`

	// The URI used to perform a Redfish Specification-defined Multipart HTTP or HTTPS push update to the update service.
	MultipartHttpPushUri string `json:"MultipartHttpPushUri,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to a collection of server certificates for the server referenced by the ImageURI property in SimpleUpdate.
	RemoteServerCertificates CertificateCollection `json:"RemoteServerCertificates,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// An inventory of software.
	SoftwareInventory SoftwareInventoryCollection `json:"SoftwareInventory,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// An indication of whether the service will verify the certificate of the server referenced by the ImageURI property in SimpleUpdate prior to sending the transfer request.
	VerifyRemoteServerCertificate bool `json:"VerifyRemoteServerCertificate,omitempty"`

}
