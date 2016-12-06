package crypto

func padding(array []byte, size int) []byte {
	pad := append(make([]byte, size), array...)
	return pad[len(pad) - size:]
}
