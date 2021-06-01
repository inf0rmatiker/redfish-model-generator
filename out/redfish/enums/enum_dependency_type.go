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
	// A simple mapping dependency. The attribute value or state is changed to the mapped value if the condition evaluates to true.
	DependencyType_MAP DependencyType = "Map"

)
