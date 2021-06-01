/* -----------------------------------------------------------------
* file_share.go -
*
* SNIA Swordfish FileShare resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// FileShare - An instance of a shared set of files with a common directory structure.
// Modeled after SNIA Swordfish schema FileShare
type FileShare struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// Continuous Availability is supported. Client/Server mediated recovery from network and server failure with application transparency.
	CASupported bool `json:"CASupported,omitempty"`

	// An array of default access capabilities for the file share. The types of default access can include Read, Write, and/or Execute.
	DefaultAccessCapabilities []StorageAccessCapability `json:"DefaultAccessCapabilities,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// A link to the collection of Ethernet interfaces that provide access to this file share.
	EthernetInterfaces map[string]interface{} `json:"EthernetInterfaces,omitempty"`

	// Execute access is supported by the file share.
	ExecuteSupport bool `json:"ExecuteSupport,omitempty"`

	// A path to an exported file or directory on the file system where this file share is hosted.
	FileSharePath string `json:"FileSharePath,omitempty"`

	// Specifies the type of quota enforcement.
	FileShareQuotaType QuotaType `json:"FileShareQuotaType,omitempty"`

	// The number of remaining bytes that may be used by this file share.
	FileShareRemainingQuotaBytes int `json:"FileShareRemainingQuotaBytes,omitempty"`

	// The maximum number of bytes that may be used by this file share.
	FileShareTotalQuotaBytes int `json:"FileShareTotalQuotaBytes,omitempty"`

	// An array of file sharing protocols supported by this file share.
	FileSharingProtocols []FileProtocol `json:"FileSharingProtocols,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links object contains the links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// An array of low space warning threshold percentages for the file share.
	LowSpaceWarningThresholdPercents []int `json:"LowSpaceWarningThresholdPercents,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Root access is allowed by the file share.
	RootAccess bool `json:"RootAccess,omitempty"`

	// Indicates the status of the file share.
	Status map[string]interface{} `json:"Status,omitempty"`

	// Defines how writes are replicated to the shared source.
	WritePolicy map[string]interface{} `json:"WritePolicy,omitempty"`

}
