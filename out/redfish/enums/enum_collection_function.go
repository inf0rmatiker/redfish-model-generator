/* -----------------------------------------------------------------
* enum_collection_function.go -
*
* DMTF Redfish CollectionFunction enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type CollectionFunction string

const (
	// An averaging function.
	CollectionFunction_AVERAGE CollectionFunction = "Average"

	// A maximum function.
	CollectionFunction_MAXIMUM CollectionFunction = "Maximum"

	// A minimum function.
	CollectionFunction_MINIMUM CollectionFunction = "Minimum"

	// A summation function.
	CollectionFunction_SUMMATION CollectionFunction = "Summation"

)
