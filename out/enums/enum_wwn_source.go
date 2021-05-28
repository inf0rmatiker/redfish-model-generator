/* -----------------------------------------------------------------
 * enum_wwn_source.go -
 *
 * DMTF Redfish WWNSource enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type WWNSource string

const (
	// The set of FC/FCoE boot targets was applied locally through API or UI.
	WWNSource_CONFIGURED_LOCALLY WWNSource = "ConfiguredLocally"

	// The set of FC/FCoE boot targets was applied by the Fibre Channel fabric.
	WWNSource_PROVIDED_BY_FABRIC WWNSource = "ProvidedByFabric"

)
