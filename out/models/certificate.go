/* -----------------------------------------------------------------
* certificate.go -
*
* DMTF Redfish Certificate resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Certificate schema describes a certificate that proves the identify of a component, account, or service.
type Certificate struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The string for the certificate.
	CertificateString string `json:"CertificateString,omitempty"`

	// The format of the certificate.
	CertificateType CertificateType `json:"CertificateType,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The fingerprint of the certificate.
	Fingerprint string `json:"Fingerprint,omitempty"`

	// The hash algorithm for the fingerprint of the certificate.
	FingerprintHashAlgorithm string `json:"FingerprintHashAlgorithm,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The issuer of the certificate.
	Issuer Identifier `json:"Issuer,omitempty"`

	// The key usage extension, which defines the purpose of the public keys in this certificate.
	KeyUsage array `json:"KeyUsage,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The serial number of the certificate.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The algorithm used for creating the signature of the certificate.
	SignatureAlgorithm string `json:"SignatureAlgorithm,omitempty"`

	// The subject of the certificate.
	Subject Identifier `json:"Subject,omitempty"`

	// The UEFI signature owner for this certificate.
	UefiSignatureOwner string `json:"UefiSignatureOwner,omitempty"`

	// The date when the certificate is no longer valid.
	ValidNotAfter string `json:"ValidNotAfter,omitempty"`

	// The date when the certificate becomes valid.
	ValidNotBefore string `json:"ValidNotBefore,omitempty"`

}
