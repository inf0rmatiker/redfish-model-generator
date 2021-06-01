/* -----------------------------------------------------------------
* enum_authentication_type.go -
*
* SNIA Swordfish AuthenticationType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AuthenticationType string

const (
	// No authentication.
	AuthenticationType_NONE AuthenticationType = "None"

	// Public Key Infrastructure.
	AuthenticationType_PKI AuthenticationType = "PKI"

	// Ticket-based (e.g., Kerberos).
	AuthenticationType_TICKET AuthenticationType = "Ticket"

	// Password/shared-secret.
	AuthenticationType_PASSWORD AuthenticationType = "Password"

)
