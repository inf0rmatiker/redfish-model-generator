/* -----------------------------------------------------------------
* stanza.go -
*
* DMTF Redfish Stanza resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A stanza contains properties that describe a request to be fulfilled within a manifest.
type Stanza struct {
	// The OEM-defined type of stanza.
	OEMStanzaType string `json:"OEMStanzaType,omitempty"`

	// The request details for the stanza.
	Request Request `json:"Request,omitempty"`

	// The response details for the stanza.
	Response Response `json:"Response,omitempty"`

	// The identifier of the stanza.  This is a unique identifier specified by the client and is not used by the service.
	StanzaId string `json:"StanzaId,omitempty"`

	// The type of stanza.
	StanzaType StanzaType `json:"StanzaType,omitempty"`

}
