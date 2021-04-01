package util

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log(os.Getenv("E_CLOUD_") == "")
}
