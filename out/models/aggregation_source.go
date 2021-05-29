/* -----------------------------------------------------------------
* aggregation_source.go -
*
* DMTF Redfish AggregationSource resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The AggregationSource schema is used to represent the source of information for a subset of the resources provided by a Redfish service.  It can be thought of as a provider of information.  As such, most such interfaces have requirements to support the gathering of information like address and account used to access the information.
type AggregationSource struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The URI of the system to be accessed.
	HostName string `json:"HostName"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The password for accessing the aggregation source.  The value is `null` in responses.
	Password string `json:"Password,omitempty"`

	// SNMP settings of the aggregation source.
	SNMP SNMPSettings `json:"SNMP,omitempty"`

	// The user name for accessing the aggregation source.
	UserName string `json:"UserName,omitempty"`

}
