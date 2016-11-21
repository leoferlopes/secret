package crypto

import "fmt"

func RSACipher(message uint16, key RSAKey) uint16 {
	fmt.Printf("%+v\n", key)
	return binExp(message, key.E, key.N)
}

func binExp(b uint16, e int, n uint16) uint16 {
	res := uint(b)
	y := uint(1)

	/* Caso base. */
	if e == 0 {
		return 1
	}

	for e > 1 {
		if e & 1 != 0 {
			/*
			 * Caso especial: expoente é ímpar.
			 * Acumular uma potência de 'res' em 'y'.
			 */
			y = (y * res) % uint(n)

			e = e - 1
		}

		/*
		 * Elevamos 'res' ao quadrado, dividimos expoente por 2.
		 */
		res = (res * res) % uint(n)

		e = e / 2
	}

	return uint16((res * y) % uint(n))
}

type RSAKey struct {
	N uint16
	E int
}