/* -----------------------------------------------------------------
* serial_interface.go -
*
* DMTF Redfish SerialInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The SerialInterface schema describes an asynchronous serial interface, such as an RS-232 interface, available to a system or device.
type SerialInterface struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The receive and transmit rate of data flow, typically in bits per second (bit/s), over the serial connection.
	BitRate BitRate `json:"BitRate,omitempty"`

	// The type of connector used for this interface.
	ConnectorType ConnectorType `json:"ConnectorType,omitempty"`

	// The number of data bits that follow the start bit over the serial connection.
	DataBits DataBits `json:"DataBits,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The type of flow control, if any, that is imposed on the serial connection.
	FlowControl FlowControl `json:"FlowControl,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether this interface is enabled.
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The type of parity used by the sender and receiver to detect errors over the serial connection.
	Parity Parity `json:"Parity,omitempty"`

	// The physical pinout configuration for a serial connector.
	PinOut PinOut `json:"PinOut,omitempty"`

	// The type of signal used for the communication connection.
	SignalType SignalType `json:"SignalType,omitempty"`

	// The period of time before the next start bit is transmitted.
	StopBits StopBits `json:"StopBits,omitempty"`

}
