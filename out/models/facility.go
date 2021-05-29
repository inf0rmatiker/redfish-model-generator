/* -----------------------------------------------------------------
* facility.go -
*
* DMTF Redfish Facility resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Facility schema represents the physical location containing equipment, such as a room, building, or campus.
type Facility struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the ambient environment metrics for this facility.
	AmbientMetrics EnvironmentMetrics `json:"AmbientMetrics,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to the environment metrics for this facility.
	EnvironmentMetrics EnvironmentMetrics `json:"EnvironmentMetrics,omitempty"`

	// The type of location this resource represents.
	FacilityType FacilityType `json:"FacilityType"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the facility.
	Location Location `json:"Location,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Link to the power domains in this facility.
	PowerDomains PowerDomainCollection `json:"PowerDomains,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
