package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func threaded_scan(data *[][]string, lock *sync.Mutex, ipuri *[]string, wg *sync.WaitGroup) {

	var ip string

	for len(*ipuri) > 0 {

		lock.Lock()

		ip = (*ipuri)[0]
		*ipuri = (*ipuri)[1:]
		fmt.Println("[RAMASE] ", len(*ipuri))

		lock.Unlock()

		scan(data, ip)
	}

	wg.Done()

}

func main() {

	//initialize
	var wg sync.WaitGroup
	var lock_ipuri sync.Mutex
	ip_list := read_ip_list()

	threads, _ := strconv.Atoi(os.Args[1:][0])
	wg.Add(threads)

	data := read_passwords()

	for i := 0; i < threads; i++ {
		go threaded_scan(&data, &lock_ipuri, &ip_list, &wg)
	}

	fmt.Println("waiting")
	wg.Wait()
}
