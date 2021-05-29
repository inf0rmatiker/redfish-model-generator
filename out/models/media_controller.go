/* -----------------------------------------------------------------
* media_controller.go -
*
* DMTF Redfish MediaController resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The MediaController schema contains the definition of the media controller and its configuration.
type MediaController struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to the environment metrics for this media controller.
	EnvironmentMetrics EnvironmentMetrics `json:"EnvironmentMetrics,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The manufacturer of this media controller.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The type of media controller.
	MediaControllerType MediaControllerType `json:"MediaControllerType,omitempty"`

	// The model of this media controller.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number of this media controller.
	PartNumber string `json:"PartNumber,omitempty"`

	// The link to the collection of ports associated with this media controller.
	Ports PortCollection `json:"Ports,omitempty"`

	// The serial number of this media controller.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The UUID for this media controller.
	UUID UUID `json:"UUID,omitempty"`

}