package random

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

func RandStr(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/2)))
	_, _ = rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:l]
}
