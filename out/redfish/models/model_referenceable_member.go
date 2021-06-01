/* -----------------------------------------------------------------
* referenceable_member.go -
*
* DMTF Redfish ReferenceableMember resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ReferenceableMember - This is the base type for addressable members of an array.
// Modeled after DMTF Redfish schema ReferenceableMember
type ReferenceableMember struct {
	// This is the identifier for the member within the collection.
	MemberId string `json:"MemberId,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
