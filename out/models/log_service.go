/* -----------------------------------------------------------------
* log_service.go -
*
* DMTF Redfish LogService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The LogService schema contains properties for monitoring and configuring a Log Service.
type LogService struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The current date and time, with UTC offset, that the Log Service uses to set or read time.
	DateTime string `json:"DateTime,omitempty"`

	// The UTC offset that the current DateTime property value contains in the `+HH:MM` format.
	DateTimeLocalOffset string `json:"DateTimeLocalOffset,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to the log entry collection.
	Entries LogEntryCollection `json:"Entries,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The format of the log entries.
	LogEntryType LogEntryTypes `json:"LogEntryType,omitempty"`

	// The maximum number of log entries that this service can have.
	MaxNumberOfRecords int `json:"MaxNumberOfRecords,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The overwrite policy for this service that takes place when the log is full.
	OverWritePolicy OverWritePolicy `json:"OverWritePolicy,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status Status `json:"Status,omitempty"`

	// A list of syslog message filters to be logged locally.
	SyslogFilters array `json:"SyslogFilters,omitempty"`

}
