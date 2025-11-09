package sjwt

import (
	"crypto/rand"
	"encoding/binary"
)

const (
	idLength       = 20
	idBitsPerChar  = 5
	idAlphabetMask = (1 << idBitsPerChar) - 1
	// readable 32 chars, (no 0, o, 1, i, l)
	// 0 and o are removed to avoid confusion with each other
	// 1, i, l are removed to avoid confusion with each other
	// extra g was added to fit 32 chars
	idAlphabetStr = "23456789abcdefgghjkmnpqrstuvwxyz"
)

var idAlphabet = []byte(idAlphabetStr)

// ID returns a 20 character, crypto-random identifier using a friendly alphabet.
func ID() string {
	out := make([]byte, idLength)
	var cache uint64
	var bits uint

	for i := 0; i < idLength; {
		if bits < idBitsPerChar {
			cache = randomUint64()
			bits = 64
		}

		index := cache & idAlphabetMask
		cache >>= idBitsPerChar
		bits -= idBitsPerChar

		if int(index) >= len(idAlphabet) {
			continue
		}

		out[i] = idAlphabet[index]
		i++
	}

	return string(out)
}

func randomUint64() uint64 {
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(buf[:])
}
