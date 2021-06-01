/* -----------------------------------------------------------------
* route_entry.go -
*
* DMTF Redfish RouteEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// RouteEntry - The RouteEntry schema describes the content of route entry rows.  Each route entry contains route sets that list the possible routes for the route entry.
// Modeled after DMTF Redfish schema RouteEntry
type RouteEntry struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The minimum number of hops.
	MinimumHopCount int `json:"MinimumHopCount,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The raw data of route entry rows.
	RawEntryHex string `json:"RawEntryHex,omitempty"`

	// The link to the collection of route set entries associated with this route.
	RouteSet map[string]interface{} `json:"RouteSet,omitempty"`

}
