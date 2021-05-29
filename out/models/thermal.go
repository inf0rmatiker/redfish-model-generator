/* -----------------------------------------------------------------
* thermal.go -
*
* DMTF Redfish Thermal resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Thermal schema describes temperature monitoring and thermal management subsystems, such as cooling fans, for a computer system or similar devices contained within a chassis.
type Thermal struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The set of fans for this chassis.
	Fans array `json:"Fans,omitempty"`

	Fans@odata.count count `json:"Fans@odata.count,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The redundancy information for the set of fans in this chassis.
	Redundancy array `json:"Redundancy,omitempty"`

	Redundancy@odata.count count `json:"Redundancy@odata.count,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The set of temperature sensors for this chassis.
	Temperatures array `json:"Temperatures,omitempty"`

	Temperatures@odata.count count `json:"Temperatures@odata.count,omitempty"`

}
