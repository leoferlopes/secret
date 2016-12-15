package crypto

type XOR struct{
	Key []byte
}

func NewXOR() *XOR{
	return &XOR{
		Key: []byte("maca"),
	}
}