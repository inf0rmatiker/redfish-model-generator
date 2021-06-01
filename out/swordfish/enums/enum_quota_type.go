/* -----------------------------------------------------------------
* enum_quota_type.go -
*
* SNIA Swordfish QuotaType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type QuotaType string

const (
	// Quotas are enabled but not enforced.
	QuotaType_SOFT QuotaType = "Soft"

	// Quotas are enabled and enforced.
	QuotaType_HARD QuotaType = "Hard"

)
