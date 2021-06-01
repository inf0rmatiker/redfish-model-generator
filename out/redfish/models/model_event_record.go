/* -----------------------------------------------------------------
* event_record.go -
*
* DMTF Redfish EventRecord resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// A context can be supplied at subscription time.  This property is the context value supplied by the subscriber.
	Context string `json:"Context,omitempty"`

	// The unique instance identifier of an event.
	EventId string `json:"EventId,omitempty"`

	// The time the event occurred.
	EventTimestamp string `json:"EventTimestamp,omitempty"`

	// The type of event.
	EventType map[string]interface{} `json:"EventType"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// The human-readable event message.
	Message string `json:"Message,omitempty"`

	// An array of message arguments that are substituted for the arguments in the message when looked up in the message registry.
	MessageArgs []s `json:"MessageArgs,omitempty"`

	// The key used to find the message in a message registry.
	MessageId string `json:"MessageId"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A link to the resource or object that originated the condition that caused the event to be generated.
	OriginOfCondition map[string]interface{} `json:"OriginOfCondition,omitempty"`

	// The severity of the event.
	Severity string `json:"Severity,omitempty"`

}
