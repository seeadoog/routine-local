package rlocal

import (
	"time"
	"testing"
	"fmt"
)

func TestGet(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(i int) {
			Set("key",i)
			defer RemoveAll()
			fmt.Println(Get("key"))

		}(i)
	}
	time.Sleep(1 * time.Second)

}
