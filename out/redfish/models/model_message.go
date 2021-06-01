/* -----------------------------------------------------------------
* message.go -
*
* DMTF Redfish Message resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Message - The message that the Redfish service returns.
// Modeled after DMTF Redfish schema Message
type Message struct {
	// The human-readable message, if provided.
	Message string `json:"Message,omitempty"`

	// This array of message arguments are substituted for the arguments in the message when looked up in the message registry.
	MessageArgs []s `json:"MessageArgs,omitempty"`

	// The key for this message used to find the message in a message registry.
	MessageId string `json:"MessageId"`

	// The severity of the message.
	MessageSeverity map[string]interface{} `json:"MessageSeverity,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A set of properties described by the message.
	RelatedProperties []s `json:"RelatedProperties,omitempty"`

	// Used to provide suggestions on how to resolve the situation that caused the error.
	Resolution string `json:"Resolution,omitempty"`

	// The severity of the errors.
	Severity string `json:"Severity,omitempty"`

}
