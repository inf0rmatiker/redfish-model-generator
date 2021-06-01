/* -----------------------------------------------------------------
* event_service.go -
*
* DMTF Redfish EventService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EventService - The Event Service resource contains properties for managing event subcriptions and generates the events sent to subscribers.  The resource has links to the actual collection of subscriptions (called Event Destinations).
// Modeled after DMTF Redfish schema EventService
type EventService struct {
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

	// This indicates whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// This is the number of attempts an event posting is retried before the subscription is terminated.
	DeliveryRetryAttempts float64 `json:"DeliveryRetryAttempts,omitempty"`

	// This represents the number of seconds between retry attempts for sending any given Event.
	DeliveryRetryIntervalSeconds float64 `json:"DeliveryRetryIntervalSeconds,omitempty"`

	// This is the types of Events that can be subscribed to.
	EventTypesForSubscription []EventType `json:"EventTypesForSubscription,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

	// This is a reference to a collection of Event Destination resources.
	Subscriptions map[string]interface{} `json:"Subscriptions,omitempty"`

}
