/* -----------------------------------------------------------------
 * enum_medium_type.go -
 *
 * DMTF Redfish MediumType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MediumType string

const (
	// The medium connected is copper.
	MediumType_COPPER MediumType = "Copper"

	// The medium connected is fiber optic.
	MediumType_FIBER_OPTIC MediumType = "FiberOptic"

)
