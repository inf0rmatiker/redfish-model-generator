/* -----------------------------------------------------------------
* location.go -
*
* DMTF Redfish Location resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Location - Location information for a schema file.
// Modeled after DMTF Redfish schema Location
type Location struct {
	// If the schema is hosted on the service in an archive file, this is the name of the file within the archive.
	ArchiveFile string `json:"ArchiveFile,omitempty"`

	// If the schema is hosted on the service in an archive file, this is the link to the archive file.
	ArchiveUri string `json:"ArchiveUri,omitempty"`

	// The language code for the file the schema is in.
	Language string `json:"Language,omitempty"`

	// Link to publicly available (canonical) URI for schema.
	PublicationUri string `json:"PublicationUri,omitempty"`

	// Link to locally available URI for schema.
	Uri string `json:"Uri,omitempty"`

}
