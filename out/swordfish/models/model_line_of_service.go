/* -----------------------------------------------------------------
* line_of_service.go -
*
* SNIA Swordfish LineOfService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// LineOfService - Describes a service option.
// Modeled after SNIA Swordfish schema LineOfService
type LineOfService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
