/* -----------------------------------------------------------------
* event_record.go -
*
* DMTF Redfish EventRecord resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EventRecord struct {
	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// A context can be supplied at subscription time.  This property is the context value supplied by the subscriber.
	Context string `json:"Context,omitempty"`

	// The identifier that correlates events with the same root cause.  If `0`, no other event is related to this event.
	EventGroupId int `json:"EventGroupId,omitempty"`

	// The unique instance identifier of an event.
	EventId string `json:"EventId,omitempty"`

	// The time the event occurred.
	EventTimestamp string `json:"EventTimestamp,omitempty"`

	// The type of event.
	EventType EventType `json:"EventType"`

	// The identifier for the member within the collection.
	MemberId string `json:"MemberId"`

	// The human-readable event message.
	Message string `json:"Message,omitempty"`

	// An array of message arguments that are substituted for the arguments in the message when looked up in the message registry.
	MessageArgs []string `json:"MessageArgs,omitempty"`

	// The identifier for the message.
	MessageId string `json:"MessageId"`

	// The severity of the message in this event.
	MessageSeverity Health `json:"MessageSeverity,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A link to the resource or object that originated the condition that caused the event to be generated.
	OriginOfCondition idRef `json:"OriginOfCondition,omitempty"`

	// The severity of the event.
	Severity string `json:"Severity,omitempty"`

	// Indicates this event is equivalent to a more specific event in this Event Group.
	SpecificEventExistsInGroup bool `json:"SpecificEventExistsInGroup,omitempty"`

}
