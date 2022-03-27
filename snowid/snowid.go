package snowid

// Package snowid provides a distributed snowflake unique id generator

import (
	"net"
	"time"
)

func NextId() int64 {
	return time.Now().Unix()
}

