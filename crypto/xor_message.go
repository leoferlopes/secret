package crypto

type XORMessage struct {
	CryptoMessage
	key Key
}

func (m *XORMessage) Encrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Encrypt()
	if err != nil {
		return nil, err
	}
	bytes := s.Bytes()
	size := len(bytes) + len(m.key) - 1;
	size -= (size % len(m.key));
	message := padding(bytes, size)
	for i := 0; i < len(message); {
		for j := 0; j < len(m.key); j++ {
			message[i] = message[i] ^ m.key[j]
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
	size := len(bytes) - 8 + len(m.key) - 1;
	size -= (size % len(m.key));
	actualSize := int(number().SetBytes(bytes[len(bytes) - 8:]).Int64())
	message := append([]byte{}, padding(bytes[:len(bytes) - 8], size)...)
	for i := 0; i < len(message); {
		for j := 0; j < len(m.key); j++ {
			message[i] = message[i] ^ m.key[j]
			i++
		}
	}
	s.SetBytes(message[len(message) - actualSize:])
	return s, nil
}

func NewXORMessage(message CryptoMessage, symmetricKey Key) CryptoMessage {
	return &XORMessage{
		CryptoMessage: message,
		key: symmetricKey,
	}
}

