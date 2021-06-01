/* -----------------------------------------------------------------
* security_capabilities.go -
*
* DMTF Redfish SecurityCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// SecurityCapabilities - This type contains security capabilities of a memory device.
// Modeled after DMTF Redfish schema SecurityCapabilities
type SecurityCapabilities struct {
	// An indication of whether this memory device supports the locking, or freezing, of the configuration.
	ConfigurationLockCapable bool `json:"ConfigurationLockCapable,omitempty"`

	// An indication of whether this memory device supports data locking.
	DataLockCapable bool `json:"DataLockCapable,omitempty"`

	// Maximum number of passphrases supported for this memory device.
	MaxPassphraseCount int `json:"MaxPassphraseCount,omitempty"`

	// An indication of whether the memory device is passphrase capable.
	PassphraseCapable bool `json:"PassphraseCapable,omitempty"`

	// The maximum number of incorrect passphrase attempts allowed before memory device is locked.
	PassphraseLockLimit int `json:"PassphraseLockLimit,omitempty"`

	// Security states supported by the memory device.
	SecurityStates []SecurityStates `json:"SecurityStates,omitempty"`

}
