/* -----------------------------------------------------------------
* enum_anti_virus_scan_trigger.go -
*
* SNIA Swordfish AntiVirusScanTrigger enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AntiVirusScanTrigger string

const (
	// No trigger.
	AntiVirusScanTrigger_NONE AntiVirusScanTrigger = "None"

	// Trigger on first read.
	AntiVirusScanTrigger_ON_FIRST_READ AntiVirusScanTrigger = "OnFirstRead"

	// Trigger on antivirus pattern file update.
	AntiVirusScanTrigger_ON_PATTERN_UPDATE AntiVirusScanTrigger = "OnPatternUpdate"

	// Trigger on object update.
	AntiVirusScanTrigger_ON_UPDATE AntiVirusScanTrigger = "OnUpdate"

	// Trigger on object rename.
	AntiVirusScanTrigger_ON_RENAME AntiVirusScanTrigger = "OnRename"

)
