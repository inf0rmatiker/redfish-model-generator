/* -----------------------------------------------------------------
* event_destination.go -
*
* DMTF Redfish EventDestination resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The EventDestination schema defines the target of an event subscription, including the event types and context to provide to the target in the Event payload.
type EventDestination struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to a collection of server certificates for the server referenced by the Destination property.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// A client-supplied string that is stored with the event destination subscription.
	Context string `json:"Context"`

	// The subscription delivery retry policy for events, where the subscription type is RedfishEvent.
	DeliveryRetryPolicy DeliveryRetryPolicy `json:"DeliveryRetryPolicy,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The URI of the destination event receiver.
	Destination string `json:"Destination,omitempty"`

	// The content types of the message that are sent to the EventDestination.
	EventFormatType EventFormatType `json:"EventFormatType,omitempty"`

	// The types of events that are sent to the destination.
	EventTypes array `json:"EventTypes,omitempty"`

	// An array of settings for HTTP headers, such as authorization information.  This array is null or an empty array in responses.  An empty array is the preferred return value on read operations.
	HttpHeaders array `json:"HttpHeaders,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether the events subscribed to will also include the entire resource or object referenced the OriginOfCondition property in the event payload.
	IncludeOriginOfCondition bool `json:"IncludeOriginOfCondition,omitempty"`

	// The list of MessageIds that the service sends.  If this property is absent or the array is empty, events with any MessageId are sent to the subscriber.
	MessageIds []string `json:"MessageIds,omitempty"`

	// A list of metric report definitions for which the service only sends related metric reports.  If this property is absent or the array is empty, metric reports that originate from any metric report definition are sent to the subscriber.
	MetricReportDefinitions array `json:"MetricReportDefinitions,omitempty"`

	MetricReportDefinitions@odata.count count `json:"MetricReportDefinitions@odata.count,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM-defined protocol type of the event connection.
	OEMProtocol string `json:"OEMProtocol,omitempty"`

	// The OEM-defined subscription type for events.
	OEMSubscriptionType string `json:"OEMSubscriptionType,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The array of Resources for which the service sends only related events.  If this property is absent or the array is empty, the service sends the events that originate from any Resource to the subscriber.
	OriginResources []map[string]string `json:"OriginResources,omitempty"`

	OriginResources@odata.count count `json:"OriginResources@odata.count,omitempty"`

	// The protocol type of the event connection.
	Protocol EventDestinationProtocol `json:"Protocol,omitempty"`

	// The list of the prefixes for the Message Registries that contain the MessageIds that are sent to this event destination.
	RegistryPrefixes []string `json:"RegistryPrefixes,omitempty"`

	// The list of Resource Type values (Schema names) that correspond to the OriginOfCondition.  The version and full namespace should not be specified.
	ResourceTypes []string `json:"ResourceTypes,omitempty"`

	// Settings for an SNMP event destination.
	SNMP SNMPSettings `json:"SNMP,omitempty"`

	// This property shall contain the status of the subscription.
	Status Status `json:"Status,omitempty"`

	// An indication of whether the subscription is for events in the OriginResources array and its subordinate Resources.  If `true` and the OriginResources array is specified, the subscription is for events in the OriginResources array and its subordinate Resources.  Note that Resources associated through the Links section are not considered subordinate.  If `false` and the OriginResources array is specified, the subscription shall be for events in the OriginResources array only.  If the OriginResources array is not present, this property shall have no relevance.
	SubordinateResources bool `json:"SubordinateResources,omitempty"`

	// The subscription type for events.
	SubscriptionType SubscriptionType `json:"SubscriptionType"`

	// A list of syslog message filters to send to a remote syslog server.
	SyslogFilters array `json:"SyslogFilters,omitempty"`

	// An indication of whether the service will verify the certificate of the server referenced by the Destination property prior to sending the event.
	VerifyCertificate bool `json:"VerifyCertificate,omitempty"`

}
