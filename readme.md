
#### install

```text
go get github.com/skyniu/routine-local
```

#### usage

````text
func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			Set("key",i)
			defer RemoveAll()
			fmt.Println("get current routine value:",Get("key"))

		}(i)
	}
	time.Sleep(1 * time.Second)

}
````