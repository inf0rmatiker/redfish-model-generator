/* -----------------------------------------------------------------
* identifier.go -
*
* DMTF Redfish Identifier resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The identifier information about a certificate.
type Identifier struct {
	// The city or locality of the organization of the entity.
	City string `json:"City,omitempty"`

	// The fully qualified domain name of the entity.
	CommonName string `json:"CommonName,omitempty"`

	// The country of the organization of the entity.
	Country string `json:"Country,omitempty"`

	// The email address of the contact within the organization of the entity.
	Email string `json:"Email,omitempty"`

	// The name of the organization of the entity.
	Organization string `json:"Organization,omitempty"`

	// The name of the unit or division of the organization of the entity.
	OrganizationalUnit string `json:"OrganizationalUnit,omitempty"`

	// The state, province, or region of the organization of the entity.
	State string `json:"State,omitempty"`

}
