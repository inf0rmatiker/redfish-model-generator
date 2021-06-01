/* -----------------------------------------------------------------
* links.go -
*
* SNIA Swordfish Links resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Links - The links object contains the links to other resources that are related to this resource.
// Modeled after SNIA Swordfish schema Links
type Links struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Collection of known and supported replica Classes of Service.
	SupportedReplicaOptions []ClassOfService `json:"SupportedReplicaOptions,omitempty"`

	SupportedReplicaOptionsOdataCount int `json:"SupportedReplicaOptions@odata.count,omitempty"`

}
