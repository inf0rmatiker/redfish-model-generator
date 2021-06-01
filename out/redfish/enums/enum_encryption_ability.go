/* -----------------------------------------------------------------
* enum_encryption_ability.go -
*
* DMTF Redfish EncryptionAbility enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EncryptionAbility string

const (
	// The drive is not capable of self-encryption.
	EncryptionAbility_NONE EncryptionAbility = "None"

	// The drive is capable of self-encryption per the Trusted Computing Group's Self Encrypting Drive Standard.
	EncryptionAbility_SELF_ENCRYPTING_DRIVE EncryptionAbility = "SelfEncryptingDrive"

	// The drive is capable of self-encryption through some other means.
	EncryptionAbility_OTHER EncryptionAbility = "Other"

)
