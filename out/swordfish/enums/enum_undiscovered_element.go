/* -----------------------------------------------------------------
* enum_undiscovered_element.go -
*
* SNIA Swordfish UndiscoveredElement enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type UndiscoveredElement string

const (
	// The source element is undiscovered.
	UndiscoveredElement_SOURCE_ELEMENT UndiscoveredElement = "SourceElement"

	// The replica element is undiscovered.
	UndiscoveredElement_REPLICA_ELEMENT UndiscoveredElement = "ReplicaElement"

)
