/* -----------------------------------------------------------------
* aggregation_service.go -
*
* DMTF Redfish AggregationService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// AggregationService - The AggregationService schema contains properties for managing aggregation operations, either on ad hoc combinations of resources or on defined sets of resources called aggregates.  Access points define the properties needed to access the entity being aggregated and connection methods describe the protocol or other semantics of the connection.
// Modeled after DMTF Redfish schema AggregationService
type AggregationService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the collection of aggregates associated with this service.
	Aggregates map[string]interface{} `json:"Aggregates,omitempty"`

	// The link to the collection of aggregation sources associated with this service.
	AggregationSources map[string]interface{} `json:"AggregationSources,omitempty"`

	// The link to the collection of connection methods associated with this service.
	ConnectionMethods map[string]interface{} `json:"ConnectionMethods,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether the aggregation service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
