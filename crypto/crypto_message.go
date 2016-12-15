package crypto

type CryptoMessage interface {
	Encrypt() (CryptoMessage, error)
	Decrypt() (CryptoMessage, error)
	Bytes() []byte
	SetBytes([]byte)
	String() string
}
