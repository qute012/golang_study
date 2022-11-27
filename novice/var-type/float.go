package var_type

import (
	"time"
	"math/rand"
)

func getFlaotVal() float32{
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Float32()
}