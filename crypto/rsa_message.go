package crypto

type RSAMessage struct {
	CryptoMessage
	rsa *RSA
}

func NewRSAMessage(message CryptoMessage, rsa *RSA) CryptoMessage {
	return &RSAMessage{
		CryptoMessage: message,
		rsa: rsa,
	}
}

func (m *RSAMessage) Encrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Encrypt()
	if err != nil {
		return nil, err
	}
	bytes := s.Bytes()
	data := make([]byte, 0)
	for _, b := range bytes {
		c := number().Exp(number().SetBytes([]byte{b}), number().SetBytes(m.rsa.E()), number().SetBytes(m.rsa.N()))
		data = append(data, padding(c.Bytes(), 2)...)
	}
	s.SetBytes(data)
	return s, nil
}

func (m *RSAMessage) Decrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Decrypt()
	if err != nil {
		return nil, err
	}
	bytes := s.Bytes()
	data := make([]byte, 0)
	for index := 0; index < len(bytes); index += 2 {
		c := number().Exp(number().SetBytes(bytes[index:index + 2]), number().SetBytes(m.rsa.D()), number().SetBytes(m.rsa.N()))
		if len(c.Bytes()) == 0 {
			data = append(data, 0)
		} else {
			data = append(data, c.Bytes()...)
		}
	}
	s.SetBytes(data)
	return s, nil
}