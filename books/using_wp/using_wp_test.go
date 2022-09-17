package usingwp

import "testing"

func BenchmarkMoveBooks(b *testing.B) {
	MoveBooks()
}
