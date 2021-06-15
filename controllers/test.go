package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type TestController struct {
}

func (t *TestController) Test(c *gin.Context) {
	//resp, err := http.Get("http://www.hao123.com")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	//if resp.StatusCode == 200 {
	//	fmt.Println("ok")
	//}

	//fmt.Println(FibonacciRecursion(8))

	//协程
	//Coroutine()

	//Channel
	//TestChannel()

	//WaitGroup
	TestWaitGroup()
}

func FibonacciRecursion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
	}
}

func Coroutine() {
	fmt.Println("run in main coroutine.")

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("run in child coroutine %d.\n", i)
		}(i)
	}

	//防止子协程还没有结束主协程就退出了
	time.Sleep(time.Second * 1)
}

func Convert() {
	i := strconv.Itoa(1)
	strconv.Atoi(i)
}

func TestChannel() {

	//有缓冲channel不要求发送和接收操作同步.
	//ch2 := make(chan int, 2)

	//无缓冲的channel由于没有缓冲发送和接收需要同步.
	ch := make(chan struct{})
	go func() {
		fmt.Println("start working")
		time.Sleep(time.Second * 1)
		//如果该 channel 中没有数据，就会一直阻塞等待，直到有值
		ch <- struct{}{}
	}()

	//接受 channel 的值
	<-ch
	fmt.Println("finished")
}

func TestWaitGroup() {
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.baidu.com/",
		"https://www.qq.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			response, _ := http.Get(url)
			fmt.Println(url + " " + response.Status)
		}(url)

	}
	wg.Wait()
}
