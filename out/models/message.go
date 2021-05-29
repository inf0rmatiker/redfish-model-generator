/* -----------------------------------------------------------------
* message.go -
*
* DMTF Redfish Message resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type represents how a message is defined within the message registry.
type Message struct {
	// The MessageArg descriptions, in order, used for this message.
	ArgDescriptions []string `json:"ArgDescriptions,omitempty"`

	// The MessageArg normative descriptions, in order, used for this message.
	ArgLongDescriptions []string `json:"ArgLongDescriptions,omitempty"`

	// The clearing logic associated with this message.  The properties within indicate that what messages are cleared by this message as well as under what conditions.
	ClearingLogic ClearingLogic `json:"ClearingLogic,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description"`

	// The normative language that describes this message's usage.
	LongDescription string `json:"LongDescription,omitempty"`

	// The actual message.
	Message string `json:"Message"`

	// The severity of the message.
	MessageSeverity Health `json:"MessageSeverity"`

	// The number of arguments in the message.
	NumberOfArgs int `json:"NumberOfArgs"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The MessageArg types, in order, for the message.
	ParamTypes array `json:"ParamTypes,omitempty"`

	// Used to provide suggestions on how to resolve the situation that caused the error.
	Resolution string `json:"Resolution"`

	// The severity of the message.
	Severity string `json:"Severity"`

}
