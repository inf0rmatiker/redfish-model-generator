/* -----------------------------------------------------------------
* enum_protocol.go -
*
* DMTF Redfish Protocol enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type Protocol string

const (
	// PCI Express.
	Protocol_PC_IE Protocol = "PCIe"

	// Advanced Host Controller Interface (AHCI).
	Protocol_AHCI Protocol = "AHCI"

	// Universal Host Controller Interface (UHCI).
	Protocol_UHCI Protocol = "UHCI"

	// Serial Attached SCSI.
	Protocol_SAS Protocol = "SAS"

	// Serial AT Attachment.
	Protocol_SATA Protocol = "SATA"

	// Universal Serial Bus (USB).
	Protocol_USB Protocol = "USB"

	// Non-Volatile Memory Express (NVMe).
	Protocol_NV_ME Protocol = "NVMe"

	// Fibre Channel.
	Protocol_FC Protocol = "FC"

	// Internet SCSI.
	Protocol_I_SCSI Protocol = "iSCSI"

	// Fibre Channel over Ethernet (FCoE).
	Protocol_F_CO_E Protocol = "FCoE"

	// Fibre Channel Protocol for SCSI.
	Protocol_FCP Protocol = "FCP"

	// FIbre CONnection (FICON).
	Protocol_FICON Protocol = "FICON"

	// NVMe over Fabrics.
	Protocol_NV_ME_OVER_FABRICS Protocol = "NVMeOverFabrics"

	// Server Message Block (SMB).  Also known as the Common Internet File System (CIFS).
	Protocol_SMB Protocol = "SMB"

	// Network File System (NFS) version 3.
	Protocol_NF_SV3 Protocol = "NFSv3"

	// Network File System (NFS) version 4.
	Protocol_NF_SV4 Protocol = "NFSv4"

	// Hypertext Transport Protocol (HTTP).
	Protocol_HTTP Protocol = "HTTP"

	// Hypertext Transfer Protocol Secure (HTTPS).
	Protocol_HTTPS Protocol = "HTTPS"

	// File Transfer Protocol (FTP).
	Protocol_FTP Protocol = "FTP"

	// SSH File Transfer Protocol (SFTP).
	Protocol_SFTP Protocol = "SFTP"

	// Internet Wide Area RDMA Protocol (iWARP).
	Protocol_I_WARP Protocol = "iWARP"

	// RDMA over Converged Ethernet Protocol.
	Protocol_RO_CE Protocol = "RoCE"

	// RDMA over Converged Ethernet Protocol Version 2.
	Protocol_RO_C_EV2 Protocol = "RoCEv2"

	// Inter-Integrated Circuit Bus.
	Protocol_I2_C Protocol = "I2C"

	// Transmission Control Protocol (TCP).
	Protocol_TCP Protocol = "TCP"

	// User Datagram Protocol (UDP).
	Protocol_UDP Protocol = "UDP"

	// Trivial File Transfer Protocol (TFTP).
	Protocol_TFTP Protocol = "TFTP"

	// GenZ.
	Protocol_GEN_Z Protocol = "GenZ"

	// Multiple Protocols.
	Protocol_MULTI_PROTOCOL Protocol = "MultiProtocol"

	// InfiniBand.
	Protocol_INFINI_BAND Protocol = "InfiniBand"

	// Ethernet.
	Protocol_ETHERNET Protocol = "Ethernet"

	// NVLink.
	Protocol_NV_LINK Protocol = "NVLink"

	// OEM-specific.
	Protocol_OEM Protocol = "OEM"

	// DisplayPort.
	Protocol_DISPLAY_PORT Protocol = "DisplayPort"

	// HDMI.
	Protocol_HDMI Protocol = "HDMI"

	// VGA.
	Protocol_VGA Protocol = "VGA"

	// DVI.
	Protocol_DVI Protocol = "DVI"

)
