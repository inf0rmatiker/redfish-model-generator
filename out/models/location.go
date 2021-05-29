/* -----------------------------------------------------------------
* location.go -
*
* DMTF Redfish Location resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Location information for a schema file.
type Location struct {
	// The name of the file in the archive, if the schema is hosted on the service in an archive file.
	ArchiveFile string `json:"ArchiveFile,omitempty"`

	// The link to an archive file, if the schema is hosted on the service in an archive file.
	ArchiveUri string `json:"ArchiveUri,omitempty"`

	// The language code for the schema file.
	Language string `json:"Language,omitempty"`

	// The link to publicly available (canonical) URI for schema.
	PublicationUri string `json:"PublicationUri,omitempty"`

	// The link to locally available URI for schema.
	Uri string `json:"Uri,omitempty"`

}
