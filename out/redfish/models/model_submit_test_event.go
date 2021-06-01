/* -----------------------------------------------------------------
* submit_test_event.go -
*
* DMTF Redfish SubmitTestEvent resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SubmitTestEvent - This action is used to generate a test event.
// Modeled after DMTF Redfish schema SubmitTestEvent
type SubmitTestEvent struct {
	// Friendly action name
	title string `json:"title,omitempty"`

	// Link to invoke action
	target string `json:"target,omitempty"`

}
