/* -----------------------------------------------------------------
* update_service.go -
*
* DMTF Redfish UpdateService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// UpdateService - The UpdateService schema describes the Update Service and the properties for the Service itself with links to collections of firmware and software inventory.  The Update Service also provides methods for updating software and firmware of the Resources in a Redfish Service.
// Modeled after DMTF Redfish schema UpdateService
type UpdateService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An inventory of firmware.
	FirmwareInventory map[string]interface{} `json:"FirmwareInventory,omitempty"`

	// The URI used to perform an HTTP or HTTPS push update to the Update Service.  The format of the message is vendor-specific.
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

	// The maximum size in bytes of the software update image that this Service supports.
	MaxImageSizeBytes int `json:"MaxImageSizeBytes,omitempty"`

	// The URI used to perform a Redfish Specification-defined Multipart HTTP or HTTPS push update to the Update Service.
	MultipartHttpPushUri string `json:"MultipartHttpPushUri,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// An inventory of software.
	SoftwareInventory map[string]interface{} `json:"SoftwareInventory,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
