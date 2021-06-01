/* -----------------------------------------------------------------
* class_of_service.go -
*
* SNIA Swordfish ClassOfService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ClassOfService - A service option composed of one or more service options.
// Modeled after SNIA Swordfish schema ClassOfService
type ClassOfService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The value identifies the current version of this class of service definition.
	ClassOfServiceVersion string `json:"ClassOfServiceVersion,omitempty"`

	// A collection of DataProtection line of service elements.
	DataProtectionLinesOfService []DataProtectionLineOfService `json:"DataProtectionLinesOfService,omitempty"`

	DataProtectionLinesOfServiceOdataCount int `json:"DataProtectionLinesOfService@odata.count,omitempty"`

	// A collection of DataSecurity line of service elements.
	DataSecurityLinesOfService []DataSecurityLineOfService `json:"DataSecurityLinesOfService,omitempty"`

	DataSecurityLinesOfServiceOdataCount int `json:"DataSecurityLinesOfService@odata.count,omitempty"`

	// A collection of DataStorage line of service elements.
	DataStorageLinesOfService []DataStorageLineOfService `json:"DataStorageLinesOfService,omitempty"`

	DataStorageLinesOfServiceOdataCount int `json:"DataStorageLinesOfService@odata.count,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// A collection of IOConnectivity line of service elements.
	IOConnectivityLinesOfService []IOConnectivityLineOfService `json:"IOConnectivityLinesOfService,omitempty"`

	IOConnectivityLinesOfServiceOdataCount int `json:"IOConnectivityLinesOfService@odata.count,omitempty"`

	// A collection of IOPerformance line of service elements.
	IOPerformanceLinesOfService []IOPerformanceLineOfService `json:"IOPerformanceLinesOfService,omitempty"`

	IOPerformanceLinesOfServiceOdataCount int `json:"IOPerformanceLinesOfService@odata.count,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
