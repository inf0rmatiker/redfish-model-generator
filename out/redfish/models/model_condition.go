/* -----------------------------------------------------------------
* condition.go -
*
* DMTF Redfish Condition resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Condition - A condition that requires attention.
// Modeled after DMTF Redfish schema Condition
type Condition struct {
	// The link to the log entry created for this condition.
	LogEntry map[string]interface{} `json:"LogEntry,omitempty"`

	// The human-readable message for this condition.
	Message string `json:"Message,omitempty"`

	// An array of message arguments that are substituted for the arguments in the message when looked up in the message registry.
	MessageArgs []s `json:"MessageArgs,omitempty"`

	// The identifier for the message.
	MessageId string `json:"MessageId"`

	// A link to the resource or object that originated the condition.
	OriginOfCondition map[string]interface{} `json:"OriginOfCondition,omitempty"`

	// The severity of the condition.
	Severity Health `json:"Severity,omitempty"`

	// The time the condition occurred.
	Timestamp string `json:"Timestamp,omitempty"`

}
