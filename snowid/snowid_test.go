package snowid

import (
	"errors"
	"testing"
	"time"
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

func TestNextIdPanics(t *testing.T) {
	t.Run("max time exceeded", func(t *testing.T) {
		defer func() { _ = recover() }()

		// SetDefaultNode()
		// epoch = time.Now().Add(time.Duration(-1*(1<<tsBits)) * time.Millisecond)
		nStartTime := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
		_ = SetNode(0, nStartTime, 41, 16, 13)

		_ = NextId()
		t.Errorf("should have panicked")
	})
}

func TestSetNode(t *testing.T) {
	SetDefaultNode()

	t.Run("invalid node id", func(t *testing.T) {
		defer func() { _ = recover() }()
		// interfaceAddrs = net.InterfaceAddrs
		nStartTime := time.Time{}
		err := SetNode(-1, nStartTime, 0, 16, 13)
		expected := errors.New("invalid node id: -ve")

		if err == nil || err.Error() != expected.Error() {
			t.Errorf("should have thown error, expected: %v; got: %v", expected, err)
		}
	})

	t.Run("invalid timestamp bits", func(t *testing.T) {
		defer func() { _ = recover() }()
		// interfaceAddrs = net.InterfaceAddrs
		nStartTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		err := SetNode(0, nStartTime, -1, 16, 13)
		expected := errors.New("invalid timestamp bits: -ve")

		if err == nil || err.Error() != expected.Error() {
			t.Errorf("should have thown error, expected: %v; got: %v", expected, err)
		}
	})

	t.Run("invalid node bits", func(t *testing.T) {
		defer func() { _ = recover() }()
		// interfaceAddrs = net.InterfaceAddrs
		nStartTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		err := SetNode(0, nStartTime, 41, -1, 13)
		expected := errors.New("invalid node bits: -ve")

		if err == nil || err.Error() != expected.Error() {
			t.Errorf("should have thown error, expected: %v; got: %v", expected, err)
		}
	})

	t.Run("invalid sequence bits", func(t *testing.T) {
		defer func() { _ = recover() }()
		// interfaceAddrs = net.InterfaceAddrs
		nStartTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		err := SetNode(0, nStartTime, 41, 16, -1)
		expected := errors.New("invalid sequence bits: -ve")

		if err == nil || err.Error() != expected.Error() {
			t.Errorf("should have thown error, expected: %v; got: %v", expected, err)
		}
	})

	t.Run("max time exceeded", func(t *testing.T) {
		defer func() { _ = recover() }()
		// interfaceAddrs = net.InterfaceAddrs
		nStartTime := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
		err := SetNode(0, nStartTime, 41, 16, 13)
		expected := errors.New("max time exceeded")

		if err == nil || err.Error() != expected.Error() {
			t.Errorf("should have thown error, expected: %v; got: %v", expected, err)
		}
	})

}

func TestParseId(t *testing.T) {
	t.Run("parse snowid", func(t *testing.T) {
		defer func() { _ = recover() }()

		nStartTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		_ = SetNode(2, nStartTime, 41, 16, 13)

		uid := NextId()
		sid := ParseId(uint64(uid))
		//t.Log(sid)
		if sid.NodeId != 2 {
			t.Errorf("expected nodeid:%v, got: %v", 2, sid.NodeId)
		}
		if sid.Timestamp > uint64(epoch.UnixMilli()) {
			t.Errorf("expected time:(%v) should post initiated  time:(%v)", sid.Timestamp, uint64(epoch.UnixMilli()))
		}
	})
}

func BenchmarkSnowid(b *testing.B) {
	SetDefaultNode()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NextId()
	}
}
