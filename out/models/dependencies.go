/* -----------------------------------------------------------------
* dependencies.go -
*
* DMTF Redfish Dependencies resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The dependency of attributes on this component.
type Dependencies struct {
	// The dependency expression for one or more attributes in this attribute registry.
	Dependency Dependency `json:"Dependency,omitempty"`

	// The AttributeName of the attribute whose change triggers the evaluation of this dependency expression.
	DependencyFor string `json:"DependencyFor,omitempty"`

	// The type of the dependency structure.
	Type DependencyType `json:"Type,omitempty"`

}
