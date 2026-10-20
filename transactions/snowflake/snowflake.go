package snowflake

import (
	"sync"
	"time"
)

var (
	snowflakeMu sync.Mutex
	lastStamp   int64 = -1
	sequence    int64 = 0
)

const (
	epoch  int64 = 1700000000000
	nodeID int64 = 1
)

func NextSnowflakeID() int64 {
	snowflakeMu.Lock()
	defer snowflakeMu.Unlock()

	now := time.Now().UnixMilli()

	if now == lastStamp {
		sequence = (sequence + 1) & 4095
		if sequence == 0 {
			for now <= lastStamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		sequence = 0
	}

	lastStamp = now

	return ((now - epoch) << 22) | (nodeID << 12) | sequence
}
