/* -----------------------------------------------------------------
* signature.go -
*
* DMTF Redfish Signature resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Signature schema describes a signature or a hash.
type Signature struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

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
	SignatureTypeRegistry SignatureTypeRegistry `json:"SignatureTypeRegistry,omitempty"`

	// The UEFI signature owner for this signature.
	UefiSignatureOwner string `json:"UefiSignatureOwner,omitempty"`

}
