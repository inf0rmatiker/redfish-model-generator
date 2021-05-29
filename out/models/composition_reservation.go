/* -----------------------------------------------------------------
* composition_reservation.go -
*
* DMTF Redfish CompositionReservation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The CompositionReservation schema contains reservation information related to the Compose action defined in the CompositionService resource when the of RequestType parameter contains the value `PreviewReserve`.
type CompositionReservation struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The client that owns the reservation.
	Client string `json:"Client,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The manifest document processed by the service that resulted in this reservation.
	Manifest Manifest `json:"Manifest,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The date time the service created the reservation.
	ReservationTime string `json:"ReservationTime,omitempty"`

	// The array of links to the reserved resource blocks.
	ReservedResourceBlocks array `json:"ReservedResourceBlocks,omitempty"`

	ReservedResourceBlocks@odata.count count `json:"ReservedResourceBlocks@odata.count,omitempty"`

}
