/* -----------------------------------------------------------------
* event_destination.go -
*
* DMTF Redfish EventDestination resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EventDestination - An Event Destination desribes the target of an event subscription, including the types of events subscribed and context to provide to the target in the Event payload.
// Modeled after DMTF Redfish schema EventDestination
type EventDestination struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// A client-supplied string that is stored with the event destination subscription.
	Context string `json:"Context"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The URI of the destination Event Service.
	Destination string `json:"Destination,omitempty"`

	// Indicates the content types of the message that will be sent to the EventDestination.
	EventFormatType map[string]interface{} `json:"EventFormatType,omitempty"`

	// This property contains the types of events that will be sent to the destination.
	EventTypes []EventType `json:"EventTypes,omitempty"`

	// This is for setting HTTP headers, such as authorization information.  This object will be null or an empty array on a GET.  An empty array is the preferred return value on GET.
	HttpHeaders []map[string]interface{} `json:"HttpHeaders,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// A list of MessageIds that the service will only send.  If this property is absent or the array is empty, then Events with any MessageId will be sent to the subscriber.
	MessageIds []string `json:"MessageIds,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A list of resources for which the service will only send related events.  If this property is absent or the array is empty, then Events originating from any resource will be sent to the subscriber.
	OriginResources []idRef `json:"OriginResources,omitempty"`

	OriginResourcesOdataCount int `json:"OriginResources@odata.count,omitempty"`

	// The protocol type of the event connection.
	Protocol EventDestinationProtocol `json:"Protocol,omitempty"`

	// A list of the Prefixes for the Message Registries that contain the MessageIds that will be sent to this event destination.
	RegistryPrefixes []string `json:"RegistryPrefixes,omitempty"`

	// A list of Resource Type values (Schema names) that correspond to the OriginOfCondition.  The version and full namespace should not be specified.
	ResourceTypes []string `json:"ResourceTypes,omitempty"`

	// By setting this to true and specifying OriginResources, this indicates the subscription will be for events from the OriginsResources specified and also all subordinate resources.  Note that resources associated via the Links section are not considered subordinate.
	SubordinateResources bool `json:"SubordinateResources,omitempty"`

	// Indicates the subscription type for events.
	SubscriptionType SubscriptionType `json:"SubscriptionType"`

}
