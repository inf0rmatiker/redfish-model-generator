/* -----------------------------------------------------------------
* clearing_logic.go -
*
* DMTF Redfish ClearingLogic resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The clearing logic associated with this message.  The properties within indicate that what messages are cleared by this message as well as under what conditions.
type ClearingLogic struct {
	// An indication of whether all prior conditions and messages are cleared, provided the ClearsIf condition is met.
	ClearsAll bool `json:"ClearsAll,omitempty"`

	// The condition when the event is cleared.
	ClearsIf ClearingType `json:"ClearsIf,omitempty"`

	// The array of MessageIds that this message clears when the other conditions are met.
	ClearsMessage []string `json:"ClearsMessage,omitempty"`

}
