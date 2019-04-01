package merkle

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestNewMerkleTree(t *testing.T) {
	txs := []string{
		"",
		"999d2c8bb6bda0bf784d9ebeb631d711dbbbfe1bc006ea13d6ad0d6a2649a971",
		"3f92594d5a3d7b4df29d7dd7c46a0dac39a96e751ba0fc9bab5435ea5e22a19d",
		"a5633f03855f541d8e60a6340fc491d49709dc821f3acb571956a856637adcb6",
		"28d97c850eaf917a4c76c02474b05b70a197eaefb468d21c22ed110afe8ec9e0",
	}

	var txsBytes [][]byte
	for _, tx := range txs {
		b, _ := hex.DecodeString(tx)
		txsBytes = append(txsBytes, b)
	}

	mt := NewMerkleTree(txsBytes)
	b, _ := hex.DecodeString("d43b669fb42cfa84695b844c0402d410213faa4f3e66cb7248f688ff19d5e5f7")
	root := hex.EncodeToString(mt.WithFirst(b))
	targetHex := "82293f182d5db07d08acf334a5a907012bbb9990851557ac0ec028116081bd5a"
	if !strings.EqualFold(root, targetHex) {
		t.Fatalf("Error cal merkle root. need: %s, got: %s", targetHex, root)
	}
}
