/* -----------------------------------------------------------------
* aggregate.go -
*
* DMTF Redfish Aggregate resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Aggregate schema describes a grouping method for an aggregation service.  Aggregates are formal groups of resources that are more persistent than ad hoc groupings.
type Aggregate struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The elements of this aggregate.
	Elements array `json:"Elements"`

	Elements@odata.count count `json:"Elements@odata.count,omitempty"`

	// The number of entries in the Elements array.
	ElementsCount int `json:"ElementsCount,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
