package dbutil

import (
	"errors"
	"go-ent-demo/types"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 固定epoch，不可修改
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 毫秒时间戳

	// ---- 关键修正：防止时间回拨 ----
	if now < w.timestamp {
		// 等待时间追上
		time.Sleep(time.Millisecond * time.Duration(w.timestamp-now))
		now = time.Now().UnixNano() / 1e6
	}

	if now == w.timestamp {
		w.number++
		if w.number > numberMax {
			// 当前毫秒序号满了，等待下一毫秒
			for now <= w.timestamp {
				time.Sleep(time.Millisecond)
				now = time.Now().UnixNano() / 1e6
			}
			w.number = 0
			w.timestamp = now
		}
	} else {
		w.number = 0
		w.timestamp = now
	}

	id := ((now - startTime) << timeShift) |
		(w.workerId << workerShift) |
		(w.number)

	return id
}

var worker *Worker

func init() {
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	worker = node
}

// func IdFunc() func() int64 {
// 	return func() int64 {
// 		id := worker.GetId()
// 		return id
// 	}
// }

func IDFunc() types.ID {
	return types.IDFrom(worker.GetId())
}
