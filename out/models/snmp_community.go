/* -----------------------------------------------------------------
* snmp_community.go -
*
* DMTF Redfish SNMPCommunity resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// An SNMP community strings.
type SNMPCommunity struct {
	// The access level of the SNMP community.
	AccessMode SNMPCommunityAccessMode `json:"AccessMode,omitempty"`

	// The SNMP community string.
	CommunityString string `json:"CommunityString,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

}
