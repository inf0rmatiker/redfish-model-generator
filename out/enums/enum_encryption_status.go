/* -----------------------------------------------------------------
 * enum_encryption_status.go -
 *
 * DMTF Redfish EncryptionStatus enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type EncryptionStatus string

const (
	// The drive is not currently encrypted.
	EncryptionStatus_UNECRYPTED EncryptionStatus = "Unecrypted"

	// The drive is currently encrypted but the data is accessible to the user in unencrypted form.
	EncryptionStatus_UNLOCKED EncryptionStatus = "Unlocked"

	// The drive is currently encrypted and the data is not accessible to the user.  However, the system can unlock the drive automatically.
	EncryptionStatus_LOCKED EncryptionStatus = "Locked"

	// The drive is currently encrypted, the data is not accessible to the user, and the system requires user intervention to expose the data.
	EncryptionStatus_FOREIGN EncryptionStatus = "Foreign"

	// The drive is not currently encrypted.
	EncryptionStatus_UNENCRYPTED EncryptionStatus = "Unencrypted"

)
