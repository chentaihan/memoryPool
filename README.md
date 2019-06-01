## memoryPool

## 内存池

### 1提供同步和非同步两种

### 2指定容量
超过容量部分会被随机删除

### 3容量排序
按照容量大小排序，使用二分查找寻找合适的内存

### 4返回内存块
1：返回的内存块大小容量 >= 需要的能存大小，返回的内存块大小 == 需要的能存大小
2：返回的内存块每个字节都初始化为0

### 5实例
```
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
	MemoryPoolTest()  
	MemoryPoolSyncTest() 
}
```