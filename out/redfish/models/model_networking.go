/* -----------------------------------------------------------------
* networking.go -
*
* DMTF Redfish Networking resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Networking - The port metrics for network ports, including Ethernet, Fibre Channel, and InfiniBand, that are not specific to one of these protocols.
// Modeled after DMTF Redfish schema Networking
type Networking struct {
	// The total number of RDMA protection errors.
	RDMAProtectionErrors int `json:"RDMAProtectionErrors,omitempty"`

	// The total number of RDMA protocol errors.
	RDMAProtocolErrors int `json:"RDMAProtocolErrors,omitempty"`

	// The total number of RDMA bytes received on a port since reset.
	RDMARXBytes int `json:"RDMARXBytes,omitempty"`

	// The total number of RDMA requests received on a port since reset.
	RDMARXRequests int `json:"RDMARXRequests,omitempty"`

	// The total number of RDMA bytes transmitted on a port since reset.
	RDMATXBytes int `json:"RDMATXBytes,omitempty"`

	// The total number of RDMA read requests transmitted on a port since reset.
	RDMATXReadRequests int `json:"RDMATXReadRequests,omitempty"`

	// The total number of RDMA requests transmitted on a port since reset.
	RDMATXRequests int `json:"RDMATXRequests,omitempty"`

	// The total number of RDMA send requests transmitted on a port since reset.
	RDMATXSendRequests int `json:"RDMATXSendRequests,omitempty"`

	// The total number of RDMA write requests transmitted on a port since reset.
	RDMATXWriteRequests int `json:"RDMATXWriteRequests,omitempty"`

	// The total number of good broadcast frames received on a port since reset.
	RXBroadcastFrames int `json:"RXBroadcastFrames,omitempty"`

	// The total number of frames discarded in a port's receive path since reset.
	RXDiscards int `json:"RXDiscards,omitempty"`

	// The total number of frames received with FCS errors on a port since reset.
	RXFCSErrors int `json:"RXFCSErrors,omitempty"`

	// The total number of false carrier errors received from phy on a port since reset.
	RXFalseCarrierErrors int `json:"RXFalseCarrierErrors,omitempty"`

	// The total number of frames received with alignment errors on a port since reset.
	RXFrameAlignmentErrors int `json:"RXFrameAlignmentErrors,omitempty"`

	// The total number of frames received on a port since reset.
	RXFrames int `json:"RXFrames,omitempty"`

	// The total number of good multicast frames received on a port since reset.
	RXMulticastFrames int `json:"RXMulticastFrames,omitempty"`

	// The total number of frames that are too long.
	RXOversizeFrames int `json:"RXOversizeFrames,omitempty"`

	// The total number of PFC frames received on a port since reset.
	RXPFCFrames int `json:"RXPFCFrames,omitempty"`

	// The total number of flow control frames from the network to pause transmission.
	RXPauseXOFFFrames int `json:"RXPauseXOFFFrames,omitempty"`

	// The total number of flow control frames from the network to resume transmission.
	RXPauseXONFrames int `json:"RXPauseXONFrames,omitempty"`

	// The total number of frames that are too short.  Short frames are less than 64 bytes.
	RXUndersizeFrames int `json:"RXUndersizeFrames,omitempty"`

	// The total number of good unicast frames received on a port since reset.
	RXUnicastFrames int `json:"RXUnicastFrames,omitempty"`

	// The total number of good broadcast frames transmitted on a port since reset.
	TXBroadcastFrames int `json:"TXBroadcastFrames,omitempty"`

	// The total number of frames discarded in a port's transmit path since reset.
	TXDiscards int `json:"TXDiscards,omitempty"`

	// The times a single transmitted frame encountered more than 15 collisions.
	TXExcessiveCollisions int `json:"TXExcessiveCollisions,omitempty"`

	// The total number of frames transmitted on a port since reset.
	TXFrames int `json:"TXFrames,omitempty"`

	// The total number of collisions that occurred after one slot time as defined by IEEE 802.3.
	TXLateCollisions int `json:"TXLateCollisions,omitempty"`

	// The total number of good multicast frames transmitted on a port since reset.
	TXMulticastFrames int `json:"TXMulticastFrames,omitempty"`

	// The times that a transmitted frame encountered 2-15 collisions.
	TXMultipleCollisions int `json:"TXMultipleCollisions,omitempty"`

	// The total number of PFC frames sent on a port since reset.
	TXPFCFrames int `json:"TXPFCFrames,omitempty"`

	// The total number of XOFF frames transmitted to the network.
	TXPauseXOFFFrames int `json:"TXPauseXOFFFrames,omitempty"`

	// The total number of XON frames transmitted to the network.
	TXPauseXONFrames int `json:"TXPauseXONFrames,omitempty"`

	// The times that a successfully transmitted frame encountered a single collision.
	TXSingleCollisions int `json:"TXSingleCollisions,omitempty"`

	// The total number of good unicast frames transmitted on a port since reset.
	TXUnicastFrames int `json:"TXUnicastFrames,omitempty"`

}
