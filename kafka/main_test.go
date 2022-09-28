package kafka

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Conn = "127.0.0.1:9092"
	os.Exit(m.Run())
}
