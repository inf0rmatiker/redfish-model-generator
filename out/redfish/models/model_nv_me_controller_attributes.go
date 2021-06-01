/* -----------------------------------------------------------------
* nv_me_controller_attributes.go -
*
* DMTF Redfish NVMeControllerAttributes resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NVMeControllerAttributes - The NVMe controller attributes for a storage controller.
// Modeled after DMTF Redfish schema NVMeControllerAttributes
type NVMeControllerAttributes struct {
	// Indicates whether or not the controller supports reporting of Namespace Granularity.
	ReportsNamespaceGranularity bool `json:"ReportsNamespaceGranularity,omitempty"`

	// Indicates whether or not the controller supports reporting of a UUID list.
	ReportsUUIDList bool `json:"ReportsUUIDList,omitempty"`

	// Indicates whether or not the controller supports a 128-bit Host Identifier.
	Supports128BitHostId bool `json:"Supports128BitHostId,omitempty"`

	// Indicates whether or not the controller supports Endurance Groups.
	SupportsEnduranceGroups bool `json:"SupportsEnduranceGroups,omitempty"`

	// Indicates whether or not the controller supports exceeding Power of Non-Operational State in order to execute controller initiated background operations in a non-operational power state.
	SupportsExceedingPowerOfNonOperationalState bool `json:"SupportsExceedingPowerOfNonOperationalState,omitempty"`

	// Indicates whether or not the controller supports NVM Sets.
	SupportsNVMSets bool `json:"SupportsNVMSets,omitempty"`

	// Indicates whether or not the controller supports Predictable Latency Mode.
	SupportsPredictableLatencyMode bool `json:"SupportsPredictableLatencyMode,omitempty"`

	// Indicates whether or not the controller supports Read Recovery Levels.
	SupportsReadRecoveryLevels bool `json:"SupportsReadRecoveryLevels,omitempty"`

	// Indicates whether or not the controller supports SQ Associations.
	SupportsSQAssociations bool `json:"SupportsSQAssociations,omitempty"`

	// Indicates whether or not the controller supports restarting Keep Alive Timer if traffic is processed from an admin command or IO during a Keep Alive Timeout interval.
	SupportsTrafficBasedKeepAlive bool `json:"SupportsTrafficBasedKeepAlive,omitempty"`

}
