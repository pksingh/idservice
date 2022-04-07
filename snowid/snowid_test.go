package snowid

import (
	"testing"
)

func TestNextIds(t *testing.T) {
	SetDefaultNode()

	next1 := NextId()
	next2 := NextId()

	if next1 == next2 {
		t.Errorf("two snowids are equal: %v", next1)
	}

	if next2 < next1 {
		t.Errorf("two snowids are generated out of order, diff(next2-next1): %v", (next2 - next1))
	}

	next3 := NextId()
	if next3 < next1 {
		t.Errorf("two snowids are generated out of order, diff(next3-next1): %v", (next3 - next1))
	}
	if next3 < next2 {
		t.Errorf("two snowids are generated out of order, diff(next3-next2): %v", (next3 - next2))
	}
}

func BenchmarkSnowid(b *testing.B) {
	SetDefaultNode()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NextId()
	}
}
