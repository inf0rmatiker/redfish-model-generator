/* -----------------------------------------------------------------
* features_registry.go -
*
* SNIA Swordfish FeaturesRegistry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// FeaturesRegistry - This is the schema definition for all Features Registries.  It represents the properties for the registries themselves.  The FeatureId is formed per the Redfish specification.  It consists of the RegistryPrefix concatenated with the version concatenated with the unique identifier for the feature registry entry.
// Modeled after SNIA Swordfish schema FeaturesRegistry
type FeaturesRegistry struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The pattern property indicates that a free-form string is the unique identifier for the feature within the registry.
	Features map[string]interface{} `json:"Features"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// This is the RFC 5646 compliant language code for the registry.
	Language string `json:"Language"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the organization or company that publishes this registry.
	OwningEntity string `json:"OwningEntity"`

	// This is the single word prefix used to form a Feature ID structure.
	RegistryPrefix string `json:"RegistryPrefix"`

	// This is the feature registry version which is used in the middle portion of a Feature ID.
	RegistryVersion string `json:"RegistryVersion"`

}
