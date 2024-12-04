package lib

import (
	"testing"
)

func TestBytes2Int(t *testing.T) {
	result := BytesToInt16([]byte{0x00, 0x01,  0x00})
	t.Log(result)
}

func TestInt2Bytes(t *testing.T) {
	result := IntToBytes(256, 3)
	t.Logf("%x", result)
}