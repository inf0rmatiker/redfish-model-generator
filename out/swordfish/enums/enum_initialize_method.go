/* -----------------------------------------------------------------
* enum_initialize_method.go -
*
* SNIA Swordfish InitializeMethod enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type InitializeMethod string

const (
	// The volume will be available for use immediately, with no preparation.
	InitializeMethod_SKIP InitializeMethod = "Skip"

	// The volume will be available for use immediately, with data erasure and preparation to happen as background tasks.
	InitializeMethod_BACKGROUND InitializeMethod = "Background"

	// Data erasure and preparation tasks will complete before the volume is presented as available for use.
	InitializeMethod_FOREGROUND InitializeMethod = "Foreground"

)
