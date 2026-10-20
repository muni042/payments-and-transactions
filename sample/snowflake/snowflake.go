package snowflake
import (
	"errors"
	"sync"
	"time"
)

const epoch int64 = 1700000000000 // custom epoch

var (
	mu        sync.Mutex
	lastStamp int64
	sequence  int64
	nodeID    int64 = 1
)

func NextSnowflakeID() int64 {
	mu.Lock()
	defer mu.Unlock()

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

	if nodeID > 31 {
		panic(errors.New("nodeID must be less than 32"))
	}

	id := ((now - epoch) << 22) | (nodeID << 12) | sequence
	return id
}
