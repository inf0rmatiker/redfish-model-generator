/* -----------------------------------------------------------------
* enum_authentication_method.go -
*
* DMTF Redfish AuthenticationMethod enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AuthenticationMethod string

const (
	// No iSCSI authentication is used.
	AuthenticationMethod_NONE AuthenticationMethod = "None"

	// iSCSI Challenge Handshake Authentication Protocol (CHAP) authentication is used.
	AuthenticationMethod_CHAP AuthenticationMethod = "CHAP"

	// iSCSI Mutual Challenge Handshake Authentication Protocol (CHAP) authentication is used.
	AuthenticationMethod_MUTUAL_CHAP AuthenticationMethod = "MutualCHAP"

)
