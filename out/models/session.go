/* -----------------------------------------------------------------
* session.go -
*
* DMTF Redfish Session resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Session Resource describes a single connection (session) between a client and a Redfish Service instance.
type Session struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The IP address of the client that created the session.
	ClientOriginIPAddress string `json:"ClientOriginIPAddress,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The active OEM-defined session type.
	OemSessionType string `json:"OemSessionType,omitempty"`

	// The password for this session.  The value is `null` in responses.
	Password string `json:"Password,omitempty"`

	// The active session type.
	SessionType SessionTypes `json:"SessionType,omitempty"`

	// The UserName for the account for this session.
	UserName string `json:"UserName,omitempty"`

}
