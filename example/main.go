package main

import (
	"fmt"
	"github.com/chentaihan/memoryPool"
	"math/rand"
)

func MemoryPoolTest() {
	//实例化内存池
	pool := memoryPool.NewMemoryPool(10)
	list := make([][]byte, 0)
	for i := 10; i < 40; i++ {
		//从内存池中获取内存
		buffer := pool.Get(i)
		list = append(list, buffer)
	}

	for _, buffer := range list {
		//将内存块存入内存池
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
	//获取指定位置的内存块
	buffer, _ = pool.GetIndex(index)
	fmt.Println(len(buffer), cap(buffer))

	for pool.Len() > 0 {
		//随机获取一个内存块
		buffer, _ := pool.GetRandom()
		fmt.Println(len(buffer), cap(buffer))
	}
	fmt.Println(pool.Len(), pool.Cap())
}

func MemoryPoolSyncTest() {
	//实例化同步内存池
	pool := memoryPool.NewMemoryPoolSync(10)
	list := make([][]byte, 0)
	for i := 10; i < 40; i++ {
		//从内存池中获取内存
		buffer := pool.Get(i)
		list = append(list, buffer)
	}

	for _, buffer := range list {
		//将内存块存入内存池
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
	//获取指定位置的内存块
	buffer, _ = pool.GetIndex(index)
	fmt.Println(len(buffer), cap(buffer))

	for pool.Len() > 0 {
		//随机获取一个内存块
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
