/* -----------------------------------------------------------------
* supported_feature.go -
*
* SNIA Swordfish SupportedFeature resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SupportedFeature - This type shall describe how a supported feature is defined within the registry.
// Modeled after SNIA Swordfish schema SupportedFeature
type SupportedFeature struct {
	// The profile definition that defines the feature.
	CorrespondingProfileDefinition string `json:"CorrespondingProfileDefinition"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description"`

	// The Name of the feature.
	FeatureName string `json:"FeatureName"`

	// The Version of the feature.
	Version string `json:"Version"`

}
