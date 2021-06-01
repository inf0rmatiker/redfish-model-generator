/* -----------------------------------------------------------------
* manifest.go -
*
* DMTF Redfish Manifest resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Manifest - This type describes a manifest containing a set of requests to be fulfilled.  The manifest contains a set of stanzas, where each stanza describes a single request.
// Modeled after DMTF Redfish schema Manifest
type Manifest struct {
	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The expansion control for references in manifest responses, similar to the `$expand=.` query parameter.
	Expand Expand `json:"Expand,omitempty"`

	// An array of stanzas that describe the requests specified by this manifest.
	Stanzas []Stanza `json:"Stanzas,omitempty"`

	// The date and time when the manifest was created.
	Timestamp string `json:"Timestamp,omitempty"`

}
