/* -----------------------------------------------------------------
 * enum_boot_order_types.go -
 *
 * DMTF Redfish BootOrderTypes enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type BootOrderTypes string

const (
	// The system uses the BootOrder property to specify the persistent boot order.
	BootOrderTypes_BOOT_ORDER BootOrderTypes = "BootOrder"

	// The system uses the AliasBootOrder property to specify the persistent boot order.
	BootOrderTypes_ALIAS_BOOT_ORDER BootOrderTypes = "AliasBootOrder"

)
