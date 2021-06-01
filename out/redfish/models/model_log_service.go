/* -----------------------------------------------------------------
* log_service.go -
*
* DMTF Redfish LogService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// LogService - This resource represents the log service for the resource or service to which it is associated.
// Modeled after DMTF Redfish schema LogService
type LogService struct {
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

	// This indicates whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The maximum number of log entries this service can have.
	MaxNumberOfRecords float64 `json:"MaxNumberOfRecords,omitempty"`

	// The overwrite policy for this service that takes place when the log is full.
	OverWritePolicy OverWritePolicy `json:"OverWritePolicy,omitempty"`

	// The current DateTime (with offset) for the log service, used to set or read time.
	DateTime string `json:"DateTime,omitempty"`

	// The time offset from UTC that the DateTime property is set to in format: +06:00 .
	DateTimeLocalOffset string `json:"DateTimeLocalOffset,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

	// References to the log entry collection.
	Entries map[string]interface{} `json:"Entries,omitempty"`

}
