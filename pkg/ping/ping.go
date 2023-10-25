package ping

import (
	"bytes"
	"encoding/binary"
	"github.com/go-ping/ping"
	"net"
	"strconv"
	"sync"
)

func ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

func back2IP(ipInt int64) string {
	a0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	a1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	a2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	a3 := strconv.FormatInt(ipInt&0xff, 10)
	addrs := a0 + "." + a1 + "." + a2 + "." + a3
	return addrs
}

func worker(addr chan string, result chan map[string]string, wg *sync.WaitGroup) {
	for a := range addr {
		pinging, err := ping.NewPinger(a)
		if err != nil {
			fping := map[string]string{
				a: "unavailable",
			}
			//fmt.Printf("%s not correct!\n", a)
			result <- fping
		}
		pinging.SetPrivileged(true)
		pinging.Count = 3
		err = pinging.Run()
		if err != nil {
			fping := map[string]string{
				a: "pong",
			}
			//fmt.Printf("%s not correct!\n", a)
			result <- fping
		} else {
			sping := map[string]string{
				a: "unavailable",
			}
			result <- sping
		}
	}
	wg.Done()
}

func Pinger(start_addrs string, end_addrs string) map[string]string {
	result := make(chan map[string]string)
	results := make(map[string]string)
	startip := ip2Long(start_addrs)
	endip := ip2Long(end_addrs)
	addrs := make(chan string, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(addrs); i++ {
		go worker(addrs, result, &wg)
	}
	for i := startip; i <= endip; i++ {
		wg.Add(1)
		i := int64(i)
		addrs <- back2IP(i)
	}
	wg.Wait()
	close(addrs)
	close(result)
	for val := range result {
		for k, v := range val {
			results[k] = v
		}
	}
	return results
}
