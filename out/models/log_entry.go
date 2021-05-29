/* -----------------------------------------------------------------
* log_entry.go -
*
* DMTF Redfish LogEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The LogEntry schema defines the record format for a log.  It is designed for Redfish event logs, OEM-specific log formats, and the IPMI System Event Log (SEL).  The EntryType field indicates the type of log and the resource includes several additional properties dependent on the EntryType.
type LogEntry struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The size of the additional data for the log entry.
	AdditionalDataSizeBytes int `json:"AdditionalDataSizeBytes,omitempty"`

	// The URI at which to access the additional data for the log entry, such as diagnostic data, image captures, or other files.
	AdditionalDataURI string `json:"AdditionalDataURI,omitempty"`

	// The date and time when the log entry was created.
	Created string `json:"Created,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The type of diagnostic data.
	DiagnosticDataType LogDiagnosticDataTypes `json:"DiagnosticDataType,omitempty"`

	// The entry code for the log entry if the entry type is `SEL`.
	EntryCode LogEntryCode `json:"EntryCode,omitempty"`

	// The type of log entry.
	EntryType LogEntryType `json:"EntryType"`

	// An identifier that correlates events with the same cause.
	EventGroupId int `json:"EventGroupId,omitempty"`

	// The unique instance identifier for an event.
	EventId string `json:"EventId,omitempty"`

	// The date and time when the event occurred.
	EventTimestamp string `json:"EventTimestamp,omitempty"`

	// The type of event recorded in this log.
	EventType EventType `json:"EventType,omitempty"`

	// An identifier of the device that has generated the IPMI SEL Event Record.
	GeneratorId string `json:"GeneratorId,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The message of the log entry.  This property decodes from the entry type.  If the entry type is `Event`, this property contains a message.  If the entry type is `SEL`, this property contains an SEL-specific message.  Otherwise, this property contains an OEM-specific log entry.  In most cases, this property contains the actual log entry.
	Message string `json:"Message,omitempty"`

	// The arguments for the message.
	MessageArgs []string `json:"MessageArgs,omitempty"`

	// The MessageId, event data, or OEM-specific information.  This property decodes from the entry type.  If the entry type is `Event`, this property contains a Redfish Specification-defined MessageId.  If the entry type is `SEL`, this property contains the Event Data.  Otherwise, this property contains OEM-specific information.
	MessageId string `json:"MessageId,omitempty"`

	// The date and time when the log entry was last modified.
	Modified string `json:"Modified,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM-defined type of diagnostic data.
	OEMDiagnosticDataType string `json:"OEMDiagnosticDataType,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The OEM-specific entry code, if the LogEntryCode type is `OEM`.
	OemLogEntryCode string `json:"OemLogEntryCode,omitempty"`

	// The OEM-specific format of the entry.  If the entry type is `Oem`, this property contains more information about the record format from the OEM.
	OemRecordFormat string `json:"OemRecordFormat,omitempty"`

	// The OEM-specific sensor type if the sensor type is `OEM`.
	OemSensorType string `json:"OemSensorType,omitempty"`

	// Used to provide suggestions on how to resolve the situation that caused the log entry.
	Resolution string `json:"Resolution,omitempty"`

	// Indicates if the cause of the log entry has been resolved or repaired.
	Resolved bool `json:"Resolved,omitempty"`

	// The IPMI-defined sensor number.
	SensorNumber int `json:"SensorNumber,omitempty"`

	// The sensor type to which the log entry pertains if the entry type is `SEL`.
	SensorType SensorType `json:"SensorType,omitempty"`

	// Indicates if the log entry has been sent to the service provider.
	ServiceProviderNotified bool `json:"ServiceProviderNotified,omitempty"`

	// The severity of the log entry.
	Severity EventSeverity `json:"Severity,omitempty"`

}
