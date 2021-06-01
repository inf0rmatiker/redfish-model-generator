/* -----------------------------------------------------------------
* enum_io_access_pattern.go -
*
* SNIA Swordfish IOAccessPattern enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type IOAccessPattern string

const (
	// Uniform distribution of reads and writes.
	IOAccessPattern_READ_WRITE IOAccessPattern = "ReadWrite"

	// Sequential read.
	IOAccessPattern_SEQUENTIAL_READ IOAccessPattern = "SequentialRead"

	// Sequential write.
	IOAccessPattern_SEQUENTIAL_WRITE IOAccessPattern = "SequentialWrite"

	// Random Read of uncached data.
	IOAccessPattern_RANDOM_READ_NEW IOAccessPattern = "RandomReadNew"

	// Random Read of cached data.
	IOAccessPattern_RANDOM_READ_AGAIN IOAccessPattern = "RandomReadAgain"

)
