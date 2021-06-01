/* -----------------------------------------------------------------
* json_schema_file.go -
*
* DMTF Redfish JsonSchemaFile resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// JsonSchemaFile - The JsonSchemaFile schema contains the properties that describe the locations, as URIs, of a Redfish Schema definition that a Redfish Service implements or references.
// Modeled after DMTF Redfish schema JsonSchemaFile
type JsonSchemaFile struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The RFC5646-conformant language codes for the available schemas.
	Languages []s `json:"Languages"`

	// Location information for this schema file.
	Location []Location `json:"Location"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The @odata.type name this schema describes.
	Schema string `json:"Schema"`

}
