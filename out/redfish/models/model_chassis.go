/* -----------------------------------------------------------------
* chassis.go -
*
* DMTF Redfish Chassis resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Chassis - A Chassis represents the physical components for any system.  This resource represents the sheet-metal confined spaces and logical zones like racks, enclosures, chassis and all other containers. Subsystems (like sensors), which operate outside of a system's data plane (meaning the resources are not accessible to software running on the system) are linked either directly or indirectly through this resource.
// Modeled after DMTF Redfish schema Chassis
type Chassis struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// This property indicates the type of physical form factor of this resource.
	ChassisType ChassisType `json:"ChassisType"`

	// This is the manufacturer of this chassis.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// This is the model number for the chassis.
	Model string `json:"Model,omitempty"`

	// This is the SKU for this chassis.
	SKU string `json:"SKU,omitempty"`

	// The serial number for this chassis.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The part number for this chassis.
	PartNumber string `json:"PartNumber,omitempty"`

	// The user assigned asset tag for this chassis.
	AssetTag string `json:"AssetTag,omitempty"`

	// The state of the indicator LED, used to identify the chassis.
	IndicatorLED IndicatorLED `json:"IndicatorLED,omitempty"`

	// Contains references to other resources that are related to this resource.
	Links map[string]interface{} `json:"Links,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	Status map[string]interface{} `json:"Status,omitempty"`

	// A reference to the logs for this chassis.
	LogServices map[string]interface{} `json:"LogServices,omitempty"`

	// A reference to the thermal properties (fans, cooling, sensors) for this chassis.
	Thermal map[string]interface{} `json:"Thermal,omitempty"`

	// A reference to the power properties (power supplies, power policies, sensors) for this chassis.
	Power map[string]interface{} `json:"Power,omitempty"`

	// This is the current power state of the chassis.
	PowerState PowerState `json:"PowerState,omitempty"`

	// The state of the physical security sensor.
	PhysicalSecurity PhysicalSecurity `json:"PhysicalSecurity,omitempty"`

	Location map[string]interface{} `json:"Location,omitempty"`

}
