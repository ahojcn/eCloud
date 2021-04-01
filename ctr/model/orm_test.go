package model

import (
	"math/rand"
	"testing"
	"time"
)

func TestGetSlave(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for true {
		time.Sleep(1 * time.Second)
		n := rand.Intn(2)
		t.Log(n)
	}
}
