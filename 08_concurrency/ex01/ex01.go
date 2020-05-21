package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	fromPort := 5200
	toPort := 5500
	var wg sync.WaitGroup
	wg.Add(toPort - fromPort)
	for i := fromPort; i <= toPort; i++ {
		go func(p int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", p))
			if err != nil {
				log.Printf("%d CLOSED (%s)\n", p, err)
				return
			}
			conn.Close()
			log.Printf("%d OPEN\n", p)
		}(i)
	}

	wg.Wait()
}
