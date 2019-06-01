package memoryPool

import (
	"math/rand"
	"sync"
)

type MemoryPoolSync struct {
	buffers []item
	sync.RWMutex
}

func NewMemoryPoolSync(poolCap int) *MemoryPoolSync {
	if poolCap < 4 {
		poolCap = 4
	}
	return &MemoryPoolSync{
		buffers: make([]item, 0, poolCap),
	}
}

func (pool *MemoryPoolSync) Get(bufferSize int) []byte {
	pool.Lock()
	defer pool.Unlock()
	if bufferSize <= 0 {
		return make([]byte, 0)
	}
	index := searchInsert(pool.buffers, len(pool.buffers), bufferSize)
	if index < 0 {
		return make([]byte, bufferSize)
	} else {
		buffer := pool.buffers[index]
		pool.buffers = append(pool.buffers[:index], pool.buffers[index+1:]...)
		return buffer[:bufferSize]
	}
}

func (pool *MemoryPoolSync) Set(buffer []byte) bool {
	if buffer == nil {
		return false
	}
	pool.reset(buffer)
	if pool.IsFull() {
		pool.GetRandom()
	}
	pool.Lock()
	defer pool.Unlock()
	if cap(pool.buffers) == len(pool.buffers) {
		return false
	}
	index := searchInsert(pool.buffers, len(pool.buffers), cap(buffer))
	pool.buffers = append(pool.buffers, buffer)
	if index >= 0 {
		for i := len(pool.buffers) - 1; i > index; i-- {
			pool.buffers[i] = pool.buffers[i-1]
		}
		pool.buffers[index] = buffer
	}
	return true
}

func (pool *MemoryPoolSync) IsFull() bool {
	pool.Lock()
	isFull := len(pool.buffers) == cap(pool.buffers)
	pool.Unlock()
	return isFull
}

func (pool *MemoryPoolSync) GetRandom() ([]byte, bool) {
	index := rand.Intn(len(pool.buffers))
	return pool.GetIndex(index)
}

func (pool *MemoryPoolSync) GetIndex(index int) ([]byte, bool) {
	pool.Lock()
	defer pool.Unlock()
	if len(pool.buffers) == 0 || index < 0 || index > len(pool.buffers) {
		return nil, false
	}
	buffer := pool.buffers[index]
	pool.buffers = append(pool.buffers[:index], pool.buffers[index+1:]...)
	return buffer, true
}

func (pool *MemoryPoolSync) Len() int {
	pool.Lock()
	size := len(pool.buffers)
	pool.Unlock()
	return size
}

func (pool *MemoryPoolSync) Cap() int {
	pool.Lock()
	capa := cap(pool.buffers)
	pool.Unlock()
	return capa
}

func (pool *MemoryPoolSync) Clear() {
	pool.Lock()
	pool.buffers = make([]item, 0, cap(pool.buffers))
	pool.Unlock()
}

func (pool *MemoryPoolSync) reset(buffer []byte) {
	i := 0
	for ; i < len(buffer); i++ {
		buffer[i] = 0
	}
	for ; i < cap(buffer); i++ {
		buffer = append(buffer, 0)
	}
}
