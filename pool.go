package memoryPool

import "math/rand"

type item []byte

type MemoryPool struct {
	buffers []item
}

func NewMemoryPool(poolCap int) *MemoryPool {
	if poolCap < 4 {
		poolCap = 4
	}
	return &MemoryPool{
		buffers: make([]item, 0, poolCap),
	}
}

func (pool *MemoryPool) Get(bufferSize int) []byte {
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

func (pool *MemoryPool) Set(buffer []byte) bool {
	if buffer == nil {
		return false
	}
	pool.reset(buffer)
	if pool.IsFull() {
		pool.GetRandom()
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

func (pool *MemoryPool) IsFull() bool {
	return len(pool.buffers) == cap(pool.buffers)
}

func (pool *MemoryPool) GetRandom() ([]byte, bool) {
	index := rand.Intn(len(pool.buffers))
	return pool.GetIndex(index)
}

func (pool *MemoryPool) GetIndex(index int) ([]byte, bool) {
	if len(pool.buffers) == 0 || index < 0 || index >= len(pool.buffers) {
		return nil, false
	}
	buffer := pool.buffers[index]
	pool.buffers = append(pool.buffers[:index], pool.buffers[index+1:]...)
	return buffer, true
}

func (pool *MemoryPool) Len() int {
	return len(pool.buffers)
}

func (pool *MemoryPool) Cap() int {
	return cap(pool.buffers)
}

func (pool *MemoryPool) Clear() {
	pool.buffers = make([]item, 0, cap(pool.buffers))
}

func (pool *MemoryPool) reset(buffer []byte) {
	i := 0
	for ; i < len(buffer); i++ {
		buffer[i] = 0
	}
	for ; i < cap(buffer); i++ {
		buffer = append(buffer, 0)
	}
}

func searchInsert(nums []item, size int, target int) int {
	if size == 0 {
		return -1
	}
	start, middle, end := 0, 0, size-1
	for start <= end {
		middle = start + (end-start)/2
		if cap(nums[middle]) == target {
			return middle
		} else if cap(nums[middle]) > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	if middle < size && target > cap(nums[middle]) {
		if target > cap(nums[size-1]) {
			return -1
		}
		return middle + 1
	}
	return middle
}
