/* -----------------------------------------------------------------
* route_set_entry.go -
*
* DMTF Redfish RouteSetEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The RouteSetEntry schema contains the information about a route.  It is part of a larger set that contains possible routes for a particular route entry.
type RouteSetEntry struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The egress interface identifier.
	EgressIdentifier int `json:"EgressIdentifier,omitempty"`

	// The number of hops.
	HopCount int `json:"HopCount,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The Virtual Channel Action index.
	VCAction int `json:"VCAction,omitempty"`

	// An indication of whether the entry is valid.
	Valid bool `json:"Valid,omitempty"`

}
