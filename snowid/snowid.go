package snowid

// Package snowid provides a distributed snowflake unique id generator

import (
	"net"
	"time"
)

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

	return nil
}

func NextId() int64 {
	return time.Now().Unix()
}

