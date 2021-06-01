/* -----------------------------------------------------------------
* event.go -
*
* DMTF Redfish Event resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Event - The Event schema describes the JSON payload received by an event destination, which has subscribed to event notification, when events occur.  This resource contains data about events, including descriptions, severity, and a message identifier to a message registry that can be accessed for further information.
// Modeled after DMTF Redfish schema Event
type Event struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// A context can be supplied at subscription time.  This property is the context value supplied by the subscriber.
	Context string `json:"Context,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Each event in this array has a set of properties that describe the event.  Because this is an array, more than one event can be sent simultaneously.
	Events []EventRecord `json:"Events"`

	EventsOdataCount int `json:"Events@odata.count,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}