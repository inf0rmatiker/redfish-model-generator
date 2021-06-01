/* -----------------------------------------------------------------
* aggregate.go -
*
* DMTF Redfish Aggregate resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Aggregate - The Aggregate schema describes a grouping method for an aggregation service.  Aggregates are formal groups of resources that are more persistent than ad hoc groupings.
// Modeled after DMTF Redfish schema Aggregate
type Aggregate struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The elements of this aggregate.
	Elements []Resource `json:"Elements"`

	ElementsOdataCount int `json:"Elements@odata.count,omitempty"`

	// The number of entries in the Elements array.
	ElementsCount int `json:"ElementsCount,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
