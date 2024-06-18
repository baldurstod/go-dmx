package dmx

import (
	"crypto/rand"
)

type DmObjectId = [16]byte

func CreateObjectId() DmObjectId {
	var b [16]byte

	_, err := rand.Read(b[:])
	if err != nil {
		panic("Can't create element id")
	}

	b[8] = (b[8] | 0x80) & 0xBF
	b[6] = (b[6] | 0x40) & 0x4F

	return b
}
