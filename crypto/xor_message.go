package crypto

type XORMessage struct {
	CryptoMessage
	XOR
}

func (m *XORMessage) Encrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Encrypt()
	if err != nil {
		return nil, err
	}
	bytes := s.Bytes()
	size := len(bytes) + len(m.Key) - 1;
	size -= (size % len(m.Key));
	message := padding(bytes, size)
	for i := 0; i < len(message); {
		for j := 0; j < len(m.Key); j++ {
			message[i] = message[i] ^ m.Key[j]
			i++
		}
	}
	s.SetBytes(append(message, padding(number(len(bytes)).Bytes(), 8)...))
	return s, nil
}

func (m *XORMessage) Decrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Decrypt()
	if err != nil {
		return nil, err
	}
	bytes := s.Bytes()
	size := len(bytes) - 8 + len(m.Key) - 1;
	size -= (size % len(m.Key));
	actualSize := int(number().SetBytes(bytes[len(bytes) - 8:]).Int64())
	message := append([]byte{}, padding(bytes[:len(bytes) - 8], size)...)
	for i := 0; i < len(message); {
		for j := 0; j < len(m.Key); j++ {
			message[i] = message[i] ^ m.Key[j]
			i++
		}
	}
	s.SetBytes(message[len(message) - actualSize:])
	return s, nil
}

func NewXORMessage(message CryptoMessage, xor *XOR) CryptoMessage {
	return &XORMessage{
		CryptoMessage: message,
		XOR: *xor,
	}
}

