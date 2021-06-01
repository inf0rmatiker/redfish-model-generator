/* -----------------------------------------------------------------
* enum_link_network_technology.go -
*
* DMTF Redfish LinkNetworkTechnology enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type LinkNetworkTechnology string

const (
	// The port is capable of connecting to an Ethernet network.
	LinkNetworkTechnology_ETHERNET LinkNetworkTechnology = "Ethernet"

	// The port is capable of connecting to an InfiniBand network.
	LinkNetworkTechnology_INFINI_BAND LinkNetworkTechnology = "InfiniBand"

	// The port is capable of connecting to a Fibre Channel network.
	LinkNetworkTechnology_FIBRE_CHANNEL LinkNetworkTechnology = "FibreChannel"

	// The port is capable of connecting to a Gen-Z fabric.
	LinkNetworkTechnology_GEN_Z LinkNetworkTechnology = "GenZ"

)
