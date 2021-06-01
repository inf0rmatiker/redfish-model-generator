/* -----------------------------------------------------------------
* log_entry.go -
*
* DMTF Redfish LogEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// LogEntry - The LogEntry schema defines the record format for a log.  It is designed for Redfish event logs, OEM-specific log formats, and the IPMI System Event Log (SEL).  The EntryType field indicates the type of log and the Resource includes several additional properties dependent on the EntryType.
// Modeled after DMTF Redfish schema LogEntry
type LogEntry struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The date and time when the log entry was created.
	Created string `json:"Created,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The entry code for the log entry if the entry type is `SEL`.
	EntryCode LogEntryCode `json:"EntryCode,omitempty"`

	// The type of log entry.
	EntryType LogEntryType `json:"EntryType"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links,omitempty"`

	// The message of the log entry.  This property decodes from the entry type.  If the entry type is `Event`, this property contains a message.  If the entry type is `SEL`, this property contains an SEL-specific message.  Otherwise, this property contains an OEM-specific log entry.  In most cases, this property contains the actual log entry.
	Message string `json:"Message,omitempty"`

	// The arguments for the message.
	MessageArgs []s `json:"MessageArgs,omitempty"`

	// The MessageId, event data, or OEM-specific information.  This property decodes from the entry type.  If the entry type is `Event`, this property contains a Redfish Specification-defined MessageId.  If the entry type is `SEL`, this property contains the Event Data.  Otherwise, this property contains OEM-specific information.
	MessageId string `json:"MessageId,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The OEM-specific format of the entry.  If the entry type is `Oem`, this property contains more information about the record format from the OEM.
	OemRecordFormat string `json:"OemRecordFormat,omitempty"`

	// The sensor number, the count of events, or OEM-specific information.  This property value is decoded from the entry type.  If the entry type is `SEL`, this property contains the sensor number.  If the entry type is `Event`, this property contains the count of events.  Otherwise, this property contains OEM-specific information.
	SensorNumber int `json:"SensorNumber,omitempty"`

	// The sensor type to which the log entry pertains if the entry type is `SEL`.
	SensorType SensorType `json:"SensorType,omitempty"`

	// The severity of the log entry.
	Severity EventSeverity `json:"Severity,omitempty"`

}
