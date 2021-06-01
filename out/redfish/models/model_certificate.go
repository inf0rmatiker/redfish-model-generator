/* -----------------------------------------------------------------
* certificate.go -
*
* DMTF Redfish Certificate resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Certificate - The Certificate schema describes a certificate that proves the identify of a component, account, or service.
// Modeled after DMTF Redfish schema Certificate
type Certificate struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The string for the certificate.
	CertificateString string `json:"CertificateString,omitempty"`

	// The format of the certificate.
	CertificateType map[string]interface{} `json:"CertificateType,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The issuer of the certificate.
	Issuer Identifier `json:"Issuer,omitempty"`

	// The key usage extension, which defines the purpose of the public keys in this certificate.
	KeyUsage []KeyUsage `json:"KeyUsage,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The subject of the certificate.
	Subject Identifier `json:"Subject,omitempty"`

	// The date when the certificate is no longer valid.
	ValidNotAfter string `json:"ValidNotAfter,omitempty"`

	// The date when the certificate becomes valid.
	ValidNotBefore string `json:"ValidNotBefore,omitempty"`

}
