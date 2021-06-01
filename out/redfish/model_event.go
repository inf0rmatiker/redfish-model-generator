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
}
