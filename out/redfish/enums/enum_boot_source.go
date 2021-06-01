/* -----------------------------------------------------------------
* enum_boot_source.go -
*
* DMTF Redfish BootSource enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type BootSource string

const (
	// Boot from the normal boot device.
	BootSource_NONE BootSource = "None"

	// Boot from the Pre-Boot EXecution (PXE) environment.
	BootSource_PXE BootSource = "Pxe"

	// Boot from the floppy disk drive.
	BootSource_FLOPPY BootSource = "Floppy"

	// Boot from the CD or DVD.
	BootSource_CD BootSource = "Cd"

	// Boot from a system BIOS-specified USB device.
	BootSource_USB BootSource = "Usb"

	// Boot from a hard drive.
	BootSource_HDD BootSource = "Hdd"

	// Boot to the BIOS setup utility.
	BootSource_BIOS_SETUP BootSource = "BiosSetup"

	// Boot to the manufacturer's utilities program or programs.
	BootSource_UTILITIES BootSource = "Utilities"

	// Boot to the manufacturer's diagnostics program.
	BootSource_DIAGS BootSource = "Diags"

	// Boot to the UEFI Shell.
	BootSource_UEFI_SHELL BootSource = "UefiShell"

	// Boot to the UEFI device specified in the UefiTargetBootSourceOverride property.
	BootSource_UEFI_TARGET BootSource = "UefiTarget"

	// Boot from an SD card.
	BootSource_SD_CARD BootSource = "SDCard"

	// Boot from a UEFI HTTP network location.
	BootSource_UEFI_HTTP BootSource = "UefiHttp"

	// Boot from a remote drive, such as an iSCSI target.
	BootSource_REMOTE_DRIVE BootSource = "RemoteDrive"

	// Boot to the UEFI device that the BootNext property specifies.
	BootSource_UEFI_BOOT_NEXT BootSource = "UefiBootNext"

)
