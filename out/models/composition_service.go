/* -----------------------------------------------------------------
* composition_service.go -
*
* DMTF Redfish CompositionService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The CompositionService schema describes a composition service and its properties and links to the resources available for composition.
type CompositionService struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the collection of resource blocks within the active pool.  Resource blocks in the active pool are contributing to at least one composed resource as a result of a composition request.
	ActivePool ResourceBlockCollection `json:"ActivePool,omitempty"`

	// An indication of whether this service is allowed to overprovision a composition relative to the composition request.
	AllowOverprovisioning bool `json:"AllowOverprovisioning,omitempty"`

	// An indication of whether a client can request that a specific resource zone fulfill a composition request.
	AllowZoneAffinity bool `json:"AllowZoneAffinity,omitempty"`

	// The link to the collection of reservations with the composition reservation collection.
	CompositionReservations CompositionReservationCollection `json:"CompositionReservations,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The link to the collection of resource blocks within the free pool.  Resource blocks in the free pool are not contributing to any composed resources.
	FreePool ResourceBlockCollection `json:"FreePool,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The length of time a composition reservation is held before the service deletes the reservation marks any related resource blocks as no longer reserved.
	ReservationDuration string `json:"ReservationDuration,omitempty"`

	// The resource blocks available on the service.
	ResourceBlocks ResourceBlockCollection `json:"ResourceBlocks,omitempty"`

	// The resource zones available on the service.
	ResourceZones ZoneCollection `json:"ResourceZones,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
