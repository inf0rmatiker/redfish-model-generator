/* -----------------------------------------------------------------
* wildcard.go -
*
* DMTF Redfish Wildcard resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The wildcard and its substitution values.
type Wildcard struct {
	// An array of values to substitute for the wildcard.
	Keys []string `json:"Keys,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name,omitempty"`

	// An array of values to substitute for the wildcard.
	Values []string `json:"Values,omitempty"`

}
