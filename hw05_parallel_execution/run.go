package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksLen := len(tasks)
	if tasksLen == 0 {
		return nil
	}

	if n < 1 {
		n = 1
	}

	var resErr error
	var errsCnt int32

	wg := sync.WaitGroup{}

	maxErrs := int32(m)
	if maxErrs < 1 {
		maxErrs = int32(tasksLen) + 1
	}

	chunkLen, chunkCnt := calcChunkParams(tasksLen, n)

	wg.Add(chunkCnt)
	for i := 0; i < chunkCnt; i++ {
		startSlice := i * chunkLen
		nextI := i + 1
		var taskChunk []Task
		if nextI < chunkCnt {
			endSlice := nextI * chunkLen
			taskChunk = tasks[startSlice:endSlice]
		} else {
			taskChunk = tasks[startSlice:]
		}

		go func(taskChunk []Task) {
			defer wg.Done()
			for _, task := range taskChunk {
				if atomic.LoadInt32(&errsCnt) >= maxErrs {
					break
				}
				err := task()
				if err != nil {
					atomic.AddInt32(&errsCnt, 1)
				}
			}
		}(taskChunk)
	}
	wg.Wait()

	if errsCnt >= maxErrs {
		resErr = ErrErrorsLimitExceeded
	}

	return resErr
}

func calcChunkParams(ttlCnt int, chunkCnt int) (chkLen int, chkCnt int) {
	if ttlCnt < chunkCnt {
		return 1, ttlCnt
	}

	chunkLen := ttlCnt / chunkCnt

	divRemainder := ttlCnt % chunkCnt
	if divRemainder > 0 {
		addChunkLen := divRemainder / chunkCnt
		chunkLen += addChunkLen
	}
	return chunkLen, chunkCnt
}
