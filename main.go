package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Названия домена: (Пример => code.org) ")
	var name string
	fmt.Scanln(&name)

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", name, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				//fmt.Printf("%d close => ", j)
				return
			}
			conn.Close()
			fmt.Printf(" open => %d", j)
		}(i)
	}
	wg.Wait()

}
