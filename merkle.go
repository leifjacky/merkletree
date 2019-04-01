/*

Ported from:
https://github.com/slush0/stratum-mining/blob/master/lib/merkletree.py
https://github.com/zone117x/node-stratum-pool/blob/master/lib/merkleTree.js

 */
package merkle

import (
	"crypto/sha256"
)

var hashFunc = func(src []byte) []byte {
	a := sha256.Sum256(src)
	b := sha256.Sum256(a[:])
	return b[:]
}

type MerkleTree struct {
	Nodes []*Node
	Steps []*Node
}

type Node struct {
	Data []byte
}

func ChangeHashFunc(f func([]byte) []byte) {
	hashFunc = f
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	t := &MerkleTree{}

	for _, nodeData := range data {
		node := &Node{
			Data: nodeData,
		}
		t.Nodes = append(t.Nodes, node)
	}

	t.calculateSteps()

	return t
}

func (t *MerkleTree) calculateSteps() {
	var steps []*Node

	PreL := []*Node{nil}
	L := t.Nodes
	Ll := len(L)

	for Ll > 1 {
		steps = append(steps, L[1])

		if Ll%2 == 1 {
			L = append(L, L[len(L)-1])
		}
		var Ld []*Node
		for i := 2; i < Ll; i += 2 {
			Ld = append(Ld, &Node{
				Data: hashFunc(append(L[i].Data, L[i+1].Data...)),
			})
		}
		L = append(PreL, Ld...)
		Ll = len(L)
	}

	t.Steps = steps
}

func (t *MerkleTree) WithFirst(hash []byte) []byte {
	for _, step := range t.Steps {
		hash = hashFunc(append(hash, step.Data...))
	}
	return hash
}

func (t *MerkleTree) MerkleRoot() []byte {
	return t.WithFirst(t.Nodes[0].Data)
}
