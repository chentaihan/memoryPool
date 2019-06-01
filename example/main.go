package main

import (
	"fmt"
	"github.com/chentaihan/memoryPool"
	"math/rand"
)

func MemoryPoolTest() {
	pool := memoryPool.NewMemoryPool(10)
	list := make([][]byte, 0)
	for i := 10; i < 40; i++ {
		buffer := pool.Get(i)
		list = append(list, buffer)
	}

	for _, buffer := range list {
		pool.Set(buffer)
	}
	fmt.Println("pool len=", pool.Len())
	buffer := pool.Get(4)
	fmt.Println("Get(4)=", len(buffer), cap(buffer))
	buffer = pool.Get(40)
	fmt.Println("Get(40)=", len(buffer), cap(buffer))
	buffer = make([]byte, 1)
	pool.Set(buffer)

	index := rand.Intn(pool.Len())
	buffer, _ = pool.GetIndex(index)
	fmt.Println(len(buffer), cap(buffer))

	for pool.Len() > 0 {
		buffer, _ := pool.GetRandom()
		fmt.Println(len(buffer), cap(buffer))
	}
	fmt.Println(pool.Len(), pool.Cap())
}

func MemoryPoolSyncTest() {
	pool := memoryPool.NewMemoryPoolSync(10)
	list := make([][]byte, 0)
	for i := 10; i < 40; i++ {
		buffer := pool.Get(i)
		list = append(list, buffer)
	}

	for _, buffer := range list {
		pool.Set(buffer)
	}
	fmt.Println("pool len=", pool.Len())
	buffer := pool.Get(4)
	fmt.Println("Get(4)=", len(buffer), cap(buffer))
	buffer = pool.Get(40)
	fmt.Println("Get(40)=", len(buffer), cap(buffer))
	buffer = make([]byte, 1)
	pool.Set(buffer)

	index := rand.Intn(pool.Len())
	buffer, _ = pool.GetIndex(index)
	fmt.Println(len(buffer), cap(buffer))

	for pool.Len() > 0 {
		buffer, _ := pool.GetRandom()
		fmt.Println(len(buffer), cap(buffer))
	}
	fmt.Println(pool.Len(), pool.Cap())
}

func main() {
	fmt.Println("MemoryPoolTest start")
	MemoryPoolTest()
	fmt.Println("MemoryPoolTest end")
	fmt.Println()
	fmt.Println("MemoryPoolSyncTest start")
	MemoryPoolSyncTest()
	fmt.Println("MemoryPoolSyncTest end")
}
