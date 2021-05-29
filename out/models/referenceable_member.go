/* -----------------------------------------------------------------
* referenceable_member.go -
*
* DMTF Redfish ReferenceableMember resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The base type for addressable members of an array.
type ReferenceableMember struct {
	OdataId string `json:"@odata.id"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
