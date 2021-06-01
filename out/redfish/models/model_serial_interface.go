/* -----------------------------------------------------------------
* serial_interface.go -
*
* DMTF Redfish SerialInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SerialInterface - This schema defines an asynchronous serial interface resource.
// Modeled after DMTF Redfish schema SerialInterface
type SerialInterface struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// This indicates whether this interface is enabled.
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	// The type of signal used for the communication connection - RS232 or RS485.
	SignalType SignalType `json:"SignalType,omitempty"`

	// The receive and transmit rate of data flow, typically in bits-per-second (bps), over the serial connection.
	BitRate BitRate `json:"BitRate,omitempty"`

	// The type of parity used by the sender and receiver in order to detect errors over the serial connection.
	Parity Parity `json:"Parity,omitempty"`

	// The number of data bits that will follow the start bit over the serial connection.
	DataBits DataBits `json:"DataBits,omitempty"`

	// The period of time before the next start bit is transmitted.
	StopBits StopBits `json:"StopBits,omitempty"`

	// The type of flow control, if any, that will be imposed on the serial connection.
	FlowControl FlowControl `json:"FlowControl,omitempty"`

	// The type of connector used for this interface.
	ConnectorType ConnectorType `json:"ConnectorType,omitempty"`

	// The physical pin configuration needed for a serial connector.
	PinOut PinOut `json:"PinOut,omitempty"`

}
