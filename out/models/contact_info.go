/* -----------------------------------------------------------------
* contact_info.go -
*
* DMTF Redfish ContactInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Contact information for this resource.
type ContactInfo struct {
	// Name of this contact.
	ContactName string `json:"ContactName,omitempty"`

	// Email address for this contact.
	EmailAddress string `json:"EmailAddress,omitempty"`

	// Phone number for this contact.
	PhoneNumber string `json:"PhoneNumber,omitempty"`

}
