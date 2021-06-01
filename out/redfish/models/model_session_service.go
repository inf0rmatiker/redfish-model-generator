/* -----------------------------------------------------------------
* session_service.go -
*
* DMTF Redfish SessionService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SessionService - This is the schema definition for the Session Service.  It represents the properties for the service itself and has links to the actual list of sessions.
// Modeled after DMTF Redfish schema SessionService
type SessionService struct {
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

	Status map[string]interface{} `json:"Status,omitempty"`

	// This indicates whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// This is the number of seconds of inactivity that a session may have before the session service closes the session due to inactivity.
	SessionTimeout float64 `json:"SessionTimeout,omitempty"`

	// Link to a collection of Sessions.
	Sessions map[string]interface{} `json:"Sessions,omitempty"`

	// The Actions object contains the available custom actions on this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

}
