/* -----------------------------------------------------------------
* signature.go -
*
* DMTF Redfish Signature resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Signature - The Signature schema describes a signature or a hash.
// Modeled after DMTF Redfish schema Signature
type Signature struct {
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

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The string for the signature.
	SignatureString string `json:"SignatureString,omitempty"`

	// The format of the signature.
	SignatureType string `json:"SignatureType,omitempty"`

	// The type of the signature.
	SignatureTypeRegistry map[string]interface{} `json:"SignatureTypeRegistry,omitempty"`

	// The UEFI signature owner for this signature.
	UefiSignatureOwner string `json:"UefiSignatureOwner,omitempty"`

}
