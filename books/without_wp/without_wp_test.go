package withoutwp

import "testing"

func BenchmarkMoveBooks(b *testing.B) {
	MoveBooks()
}
