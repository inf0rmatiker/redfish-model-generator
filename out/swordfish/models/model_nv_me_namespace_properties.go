/* -----------------------------------------------------------------
* nv_me_namespace_properties.go -
*
* SNIA Swordfish NVMeNamespaceProperties resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NVMeNamespaceProperties - This contains properties to use when Volume is used to describe an NVMe Namespace.
// Modeled after SNIA Swordfish schema NVMeNamespaceProperties
type NVMeNamespaceProperties struct {
	// The LBA data size and metadata size combination that the namespace has been formatted with.
	FormattedLBASize string `json:"FormattedLBASize,omitempty"`

	// Indicates the namespace is shareable.
	IsShareable bool `json:"IsShareable,omitempty"`

	// This property indicates whether or not the metadata is transferred at the end of the LBA creating an extended data LBA.
	MetadataTransferredAtEndOfDataLBA bool `json:"MetadataTransferredAtEndOfDataLBA,omitempty"`

	// The version of the NVMe Base Specification supported.
	NVMeVersion string `json:"NVMeVersion,omitempty"`

	// This property contains a set of Namespace Features.
	NamespaceFeatures NamespaceFeatures `json:"NamespaceFeatures,omitempty"`

	// The NVMe Namespace Identifier for this namespace.
	NamespaceId string `json:"NamespaceId,omitempty"`

	// The number of LBA data size and metadata size combinations supported by this namespace. The value of this property is between 0 and 16.
	NumberLBAFormats int `json:"NumberLBAFormats,omitempty"`

}
