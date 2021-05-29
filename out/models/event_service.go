/* -----------------------------------------------------------------
* event_service.go -
*
* DMTF Redfish EventService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The EventService schema contains properties for managing event subscriptions and generates the events sent to subscribers.  The resource has links to the actual collection of subscriptions, which are called event destinations.
type EventService struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The number of times that the POST of an event is retried before the subscription terminates.  This retry occurs at the service level, which means that the HTTP POST to the event destination fails with an HTTP `4XX` or `5XX` status code or an HTTP timeout occurs this many times before the event destination subscription terminates.
	DeliveryRetryAttempts int `json:"DeliveryRetryAttempts,omitempty"`

	// The interval, in seconds, between retry attempts for sending any event.
	DeliveryRetryIntervalSeconds int `json:"DeliveryRetryIntervalSeconds,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The content types of the message that this service can send to the event destination.
	EventFormatTypes array `json:"EventFormatTypes,omitempty"`

	// The types of events to which a client can subscribe.
	EventTypesForSubscription array `json:"EventTypesForSubscription,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether the service supports including the resource payload of the origin of condition in the event payload.
	IncludeOriginOfConditionSupported bool `json:"IncludeOriginOfConditionSupported,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The list of the prefixes of the message registries that can be used for the RegistryPrefix property on a subscription.  If this property is absent or contains an empty array, the service does not support RegistryPrefix-based subscriptions.
	RegistryPrefixes []string `json:"RegistryPrefixes,omitempty"`

	// The list of @odata.type values, or schema names, that can be specified in the ResourceTypes array in a subscription.  If this property is absent or contains an empty array, the service does not support resource type-based subscriptions.
	ResourceTypes []string `json:"ResourceTypes,omitempty"`

	// Settings for SMTP event delivery.
	SMTP SMTP `json:"SMTP,omitempty"`

	// The set of properties that are supported in the `$filter` query parameter for the ServerSentEventUri.
	SSEFilterPropertiesSupported SSEFilterPropertiesSupported `json:"SSEFilterPropertiesSupported,omitempty"`

	// The link to a URI for receiving Server-Sent Event representations for the events that this service generates.
	ServerSentEventUri string `json:"ServerSentEventUri,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// An indication of whether the service supports the SubordinateResources property on both event subscriptions and generated events.
	SubordinateResourcesSupported bool `json:"SubordinateResourcesSupported,omitempty"`

	// The link to a collection of event destinations.
	Subscriptions EventDestinationCollection `json:"Subscriptions,omitempty"`

}
