package util

import (
	"os"
)

func TransferBuffer(in *os.File, out *os.File)  {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the input into the buffer.
	for length, err := in.Read(buf); length > 0; length, err = in.Read(buf) {
		if err != nil {
			panic(err)
		}

		// Write the buffer into the output
		out.Write(buf[0:length])
	}

}
