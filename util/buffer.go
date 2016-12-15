package util

type Stream interface {
	// Read reads data from the connection.
	// Read can be made to time out and return a Error with Timeout() == true
	// after a fixed time limit; see SetDeadline and SetReadDeadline.
	Read(b []byte) (n int, err error)

	// Write writes data to the connection.
	// Write can be made to time out and return a Error with Timeout() == true
	// after a fixed time limit; see SetDeadline and SetWriteDeadline.
	Write(b []byte) (n int, err error)
}

type Cypher interface {
	Encrypt(b []byte, sequence uint64) []byte
	Decrypt(b []byte) ([]byte, uint64)
}

func TransferEncryptBuffer(in Stream, out Stream, cypher Cypher) {
	buf := make([]byte, 504)

	sequence := uint64(0)
	for length, err := in.Read(buf); length > 0; length, err = in.Read(buf) {
		println(length)
		if err != nil {
			panic(err)
		}
		bytes := cypher.Encrypt(buf[0:length], sequence)

		out.Write(bytes)
		sequence++
	}
}

func TransferDecryptBuffer(in Stream, out Stream, cypher Cypher) {
	buf := make([]byte, 1024)

	sequence := uint64(0)
	for length, err := in.Read(buf); length > 0; length, err = in.Read(buf) {
		println(length)
		if err != nil {
			panic(err)
		}
		bytes, seq := cypher.Decrypt(buf[0:length])
		if seq != sequence {
			println("seq", seq)
			println("sequence", sequence)
		}

		out.Write(bytes)
		sequence++
	}
}
