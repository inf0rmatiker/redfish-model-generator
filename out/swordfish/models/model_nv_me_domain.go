/* -----------------------------------------------------------------
* nv_me_domain.go -
*
* SNIA Swordfish NVMeDomain resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NVMeDomain - Properties for the Domain.
// Modeled after SNIA Swordfish schema NVMeDomain
type NVMeDomain struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// A collection of available firmware images.
	AvailableFirmwareImages []NVMeFirmwareImage `json:"AvailableFirmwareImages,omitempty"`

	AvailableFirmwareImagesOdataCount int `json:"AvailableFirmwareImages@odata.count,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The members of the domain.
	DomainMembers []Resource `json:"DomainMembers,omitempty"`

	DomainMembersOdataCount int `json:"DomainMembers@odata.count,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The maximum capacity per endurance group in bytes of this NVMe Domain.
	MaximumCapacityPerEnduranceGroupBytes int `json:"MaximumCapacityPerEnduranceGroupBytes,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The total capacity in bytes of this NVMe Domain.
	TotalDomainCapacityBytes int `json:"TotalDomainCapacityBytes,omitempty"`

	// The total unallocated capacity in bytes of this NVMe Domain.
	UnallocatedDomainCapacityBytes int `json:"UnallocatedDomainCapacityBytes,omitempty"`

}
