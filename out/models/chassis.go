/* -----------------------------------------------------------------
* chassis.go -
*
* DMTF Redfish Chassis resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The Chassis schema represents the physical components of a system.  This resource represents the sheet-metal confined spaces and logical zones such as racks, enclosures, chassis and all other containers.  Subsystems, such as sensors, that operate outside of a system's data plane are linked either directly or indirectly through this resource.  A subsystem that operates outside of a system's data plane are not accessible to software that runs on the system.
type Chassis struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The link to the assembly associated with this chassis.
	Assembly Assembly `json:"Assembly,omitempty"`

	// The user-assigned asset tag of this chassis.
	AssetTag string `json:"AssetTag,omitempty"`

	// The type of physical form factor of the chassis.
	ChassisType ChassisType `json:"ChassisType"`

	// The depth of the chassis.
	DepthMm float64 `json:"DepthMm,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The ASHRAE Environmental Class for this chassis.
	EnvironmentalClass EnvironmentalClass `json:"EnvironmentalClass,omitempty"`

	// The height of the chassis.
	HeightMm float64 `json:"HeightMm,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The state of the indicator LED, which identifies the chassis.
	IndicatorLED IndicatorLED `json:"IndicatorLED,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The location of the chassis.
	Location Location `json:"Location,omitempty"`

	// The link to the logs for this chassis.
	LogServices LogServiceCollection `json:"LogServices,omitempty"`

	// The manufacturer of this chassis.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The model number of the chassis.
	Model string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The link to the collection of network adapters associated with this chassis.
	NetworkAdapters NetworkAdapterCollection `json:"NetworkAdapters,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to the PCIe slot properties for this chassis.
	PCIeSlots PCIeSlots `json:"PCIeSlots,omitempty"`

	// The part number of the chassis.
	PartNumber string `json:"PartNumber,omitempty"`

	// The state of the physical security sensor.
	PhysicalSecurity PhysicalSecurity `json:"PhysicalSecurity,omitempty"`

	// The link to the power properties, or power supplies, power policies, and sensors, for this chassis.
	Power Power `json:"Power,omitempty"`

	// The current power state of the chassis.
	PowerState PowerState `json:"PowerState,omitempty"`

	// The SKU of the chassis.
	SKU string `json:"SKU,omitempty"`

	// The link to the collection of sensors located in the equipment and sub-components.
	Sensors SensorCollection `json:"Sensors,omitempty"`

	// The serial number of the chassis.
	SerialNumber string `json:"SerialNumber,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The link to the thermal properties, such as fans, cooling, and sensors, for this chassis.
	Thermal Thermal `json:"Thermal,omitempty"`

	// The UUID for this chassis.
	UUID UUID `json:"UUID,omitempty"`

	// The weight of the chassis.
	WeightKg float64 `json:"WeightKg,omitempty"`

	// The width of the chassis.
	WidthMm float64 `json:"WidthMm,omitempty"`

}
