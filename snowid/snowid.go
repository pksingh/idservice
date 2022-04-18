package snowid

// Package snowid provides a distributed snowflake unique id generator

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

// A snow ID is max 63 bits composed of
//
//	tsBits - used for time : here we keep time in ms; hence 40 bits, user configurable
//	nodeBits - used for uniq nodes [ can be combined with node with muliple pods ] : 16 bits, user configurable
//	seqBits - used for a sequence/counter number : 13 bits, user configurable
var (
	epoch   time.Time // will store unix epoch timestamp
	nodeId  int64     // will hold the current nodeid
	seq     int64     // will hold current sequece/counter value
	elapsed int64     // number of epoch passed on

	seqBits       int64 // count/sequence bits
	nodeBits      int64 // node/machine/pod id bits
	tsBits        int64 // timestamp bits
	nodeBitsShift int64 // number of bits to shift for node ID
	tsBitsShift   int64 // timestamp bits shift
	maxSeq        int64 // 2^seqBits - 1
	maxTS         int64 // 2^tsBits - 1
)

// SetNode will init nodeid, starttime, and bits used for snowid - Configure as per need
func SetNode(nId int64, nStartTime time.Time, nTimeBits, nNodeBits, nCountBits int64) error {
	if nNodeBits < 0 {
		return errors.New("invalid node bits: -ve")
	}
	nodeBits = nNodeBits

	if nCountBits < 0 {
		return errors.New("invalid sequence bits: -ve")
	}
	seqBits = nCountBits

	if nTimeBits < 0 {
		return errors.New("invalid timestamp bits: -ve")
	}
	tsBits = nTimeBits

	nodeBitsShift = seqBits                // number of bits to shift for node ID
	tsBitsShift = nodeBits + nodeBitsShift // timestamp bits shift
	maxSeq = (1 << seqBits) - 1            // 2^seqBits - 1
	maxTS = (1 << tsBits) - 1              // 2^tsBits - 1

	if nId < 0 {
		return errors.New("invalid node id: -ve")
	}
	nodeId = nId << nodeBitsShift

	if nStartTime.IsZero() {
		return errors.New("invalid start time: 0/nil")
	}
	epoch = nStartTime
	elapsed = 0

	e := time.Since(epoch).Milliseconds()
	if e > maxTS {
		return errors.New("max time exceeded")
	}
	atomic.StoreInt64(&elapsed, time.Since(epoch).Milliseconds())
	return nil
}

func SetDefaultNode() {
	nStartTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	_ = SetNode(0, nStartTime, 41, 16, 13)
}

// NextId returns the next unique snowid.
func NextId() int64 {
	t := time.Since(epoch).Milliseconds()

	// have we reach maxtime
	if t > maxTS {
		panic(fmt.Sprintf("max time exceeded: currenttime: %v, maxtime: %v", t, maxTS))
		// return errors.New("max time exceeded")
	}

	// check clock rollbacks - double check
	e := atomic.LoadInt64(&elapsed)
	for e > t {
		t = time.Since(epoch).Milliseconds()
		e = atomic.LoadInt64(&elapsed)
	}
	// have all sequence used/generated
	s := atomic.AddInt64(&seq, 1)
	for s > maxSeq {
		for t == e {
			t = time.Since(epoch).Milliseconds()
		}
		atomic.StoreInt64(&seq, 0)
		s = 0
	}
	atomic.StoreInt64(&elapsed, t)
	return (t << tsBitsShift) | nodeId | s
}

// SID snowid based on snowflake
type SnowID struct {
	Sequence  uint64
	NodeId    uint64
	Timestamp uint64
	ID        uint64
}

// ParseId parse snowid it to SID struct.
func ParseId(id uint64) SnowID {
	time := id >> (uint64(seqBits) + uint64(nodeBits))
	sequence := id & uint64(maxSeq)
	maxNodes := (1 << nodeBits) - 1
	nodeId := (id & (uint64(maxNodes) << uint64(seqBits))) >> seqBits

	return SnowID{
		ID:        id,
		Sequence:  sequence,
		NodeId:    nodeId,
		Timestamp: time,
	}
}

