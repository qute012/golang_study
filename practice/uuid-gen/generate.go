package ug

import (
	"github.com/google/uuid"
)

func Generate() string {
	u, err := uuid.NewRandom()	// v4 uuid
	if err != nil {
		panic(err)
	}
	return u.String()
}