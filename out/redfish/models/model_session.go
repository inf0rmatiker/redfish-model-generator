/* -----------------------------------------------------------------
* session.go -
*
* DMTF Redfish Session resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Session - The Session resource describes a single connection (session) between a client and a Redfish service instance.
// Modeled after DMTF Redfish schema Session
type Session struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The UserName for the account for this session.
	UserName string `json:"UserName,omitempty"`

	// This property is used in a POST to specify a password when creating a new session.  This property is null on a GET.
	Password string `json:"Password,omitempty"`

}
