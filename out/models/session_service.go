/* -----------------------------------------------------------------
* session_service.go -
*
* DMTF Redfish SessionService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The SessionService schema describes the session service and its properties, with links to the actual list of sessions.
type SessionService struct {
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

	// An indication of whether this service is enabled.  If `true`, this service is enabled.  If `false`, it is disabled, and new sessions cannot be created, old sessions cannot be deleted, and established sessions can continue operating.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The number of seconds of inactivity that a session can have before the session service closes the session due to inactivity.
	SessionTimeout int `json:"SessionTimeout,omitempty"`

	// The link to a collection of sessions.
	Sessions SessionCollection `json:"Sessions,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
