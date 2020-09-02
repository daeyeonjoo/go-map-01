package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = make(map[string]string)

	for i := 0; i < 10; i++ {
		var k = fmt.Sprintf("k%d", i)
		var v = fmt.Sprintf("v%d", i)
		m[k] = v
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(m))

	for k, v := range m {
		go func(key string, message string) {
			fmt.Println(key, message)
			waitGroup.Done()
		}(k, v)
	}

	waitGroup.Wait()
}
