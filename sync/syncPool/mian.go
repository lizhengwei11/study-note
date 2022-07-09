package syncPool

//sync.waitGroup  等待组
/*
func main() {
	sync.WaitGroup{}
}
*/

/*func main() {
	// 建立对象
	var pipe = &sync.Pool{}

	// 准备放入的字符串
	val := "Hello,World!"
	val2 := "Hello, BeiJing"
	// 放入
	pipe.Put(val2)
	pipe.Put(val)

	// 取出
	first := pipe.Get()

	// 再取就没有了,会自动调用NEW
	second := pipe.Get()
	fmt.Println(first)
	fmt.Println(second)
	s := pipe.Get()
	//取完就没了
	fmt.Println(s)
}*/

/*//互斥锁  sync.Mutex
var num int
var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go Add()
	}
	wg.Wait()
	fmt.Println(num)
}

func Add() {
	mu.Lock() // 加锁
	defer func() {
		mu.Unlock() // 解锁
		wg.Done()
	}()

	num++
}*/

/*//读写互斥锁
var wg sync.WaitGroup
var rwlock sync.RWMutex
var name string = "读"

func read() {
	defer wg.Done()
	rwlock.RLock() //读锁定  其他协程可以读  但阻止写
	fmt.Println("开始读取数据", name)
	time.Sleep(time.Second)
	fmt.Println("读取成功", name)
	rwlock.RUnlock() //读解锁
}

func write() {
	defer wg.Done()

	rwlock.Lock()
	fmt.Println("开始修改数据", name)
	name = "写入成功"
	time.Sleep(time.Second * 5)
	fmt.Println("修改成功", name)
	rwlock.Unlock()
}
func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go read()
	}
	for i := 0; i < 1; i++ {
		go write()
	}
	wg.Wait()
}
*/
