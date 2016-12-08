package util

func TransferBuffer(in stream, out stream) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the input into the buffer.
	for length, err := in.Read(buf); length > 0; length, err = in.Read(buf) {
		println(length)
		if err != nil {
			panic(err)
		}

		// Write the buffer into the output
		out.Write(buf[0:length])
	}
	println("dsd")

}

type stream interface {
	// Read reads data from the connection.
	// Read can be made to time out and return a Error with Timeout() == true
	// after a fixed time limit; see SetDeadline and SetReadDeadline.
	Read(b []byte) (n int, err error)

	// Write writes data to the connection.
	// Write can be made to time out and return a Error with Timeout() == true
	// after a fixed time limit; see SetDeadline and SetWriteDeadline.
	Write(b []byte) (n int, err error)
}
