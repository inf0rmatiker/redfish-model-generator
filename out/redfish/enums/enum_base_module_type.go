/* -----------------------------------------------------------------
* enum_base_module_type.go -
*
* DMTF Redfish BaseModuleType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type BaseModuleType string

const (
	// Registered DIMM.
	BaseModuleType_RDIMM BaseModuleType = "RDIMM"

	// UDIMM.
	BaseModuleType_UDIMM BaseModuleType = "UDIMM"

	// SO_DIMM.
	BaseModuleType_SO_DIMM BaseModuleType = "SO_DIMM"

	// Load Reduced.
	BaseModuleType_LRDIMM BaseModuleType = "LRDIMM"

	// Mini_RDIMM.
	BaseModuleType_MINI_RDIMM BaseModuleType = "Mini_RDIMM"

	// Mini_UDIMM.
	BaseModuleType_MINI_UDIMM BaseModuleType = "Mini_UDIMM"

	// SO_RDIMM_72b.
	BaseModuleType_SO_RDIMM_72B BaseModuleType = "SO_RDIMM_72b"

	// SO_UDIMM_72b.
	BaseModuleType_SO_UDIMM_72B BaseModuleType = "SO_UDIMM_72b"

	// SO_DIMM_16b.
	BaseModuleType_SO_DIMM_16B BaseModuleType = "SO_DIMM_16b"

	// SO_DIMM_32b.
	BaseModuleType_SO_DIMM_32B BaseModuleType = "SO_DIMM_32b"

	// A die within a package.
	BaseModuleType_DIE BaseModuleType = "Die"

)
