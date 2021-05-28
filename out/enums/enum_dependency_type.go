/* -----------------------------------------------------------------
 * enum_dependency_type.go -
 *
 * DMTF Redfish DependencyType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type DependencyType string

const (
	// A simple mapping dependency.  If the condition evaluates to `true`, the attribute or state changes to the mapped value.
	DependencyType_MAP DependencyType = "Map"

)
