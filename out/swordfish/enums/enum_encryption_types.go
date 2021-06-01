/* -----------------------------------------------------------------
* enum_encryption_types.go -
*
* SNIA Swordfish EncryptionTypes enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type EncryptionTypes string

const (
	// The volume is utilizing the native drive encryption capabilities of the drive hardware.
	EncryptionTypes_NATIVE_DRIVE_ENCRYPTION EncryptionTypes = "NativeDriveEncryption"

	// The volume is being encrypted by the storage controller entity.
	EncryptionTypes_CONTROLLER_ASSISTED EncryptionTypes = "ControllerAssisted"

	// The volume is being encrypted by software running on the system or the operating system.
	EncryptionTypes_SOFTWARE_ASSISTED EncryptionTypes = "SoftwareAssisted"

)
