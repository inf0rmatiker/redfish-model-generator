/* -----------------------------------------------------------------
* composition_reservation.go -
*
* DMTF Redfish CompositionReservation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CompositionReservation - The CompositionReservation schema contains reservation information related to the Compose action defined in the CompositionService resource when the of RequestType parameter contains the value `PreviewReserve`.
// Modeled after DMTF Redfish schema CompositionReservation
type CompositionReservation struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The client that owns the reservation.
	Client string `json:"Client,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The manifest document processed by the service that resulted in this reservation.
	Manifest map[string]interface{} `json:"Manifest,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The date time the service created the reservation.
	ReservationTime string `json:"ReservationTime,omitempty"`

	// The array of links to the reserved resource blocks.
	ReservedResourceBlocks []ResourceBlock `json:"ReservedResourceBlocks,omitempty"`

	ReservedResourceBlocksOdataCount int `json:"ReservedResourceBlocks@odata.count,omitempty"`

}
