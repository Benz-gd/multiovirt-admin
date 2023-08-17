package ping

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func worker(addr chan string, wg *sync.WaitGroup) {
	for a := range addr {
		pinger, err := ping.NewPinger(a)
		if err != nil {
			fmt.Printf("%s not correct!\n", a)
		}
		pinger.SetPrivileged(true)
		pinger.Count = 2
		err = pinger.Run()
		if err != nil {
			fmt.Printf("%s not connect!\n", a)
		}
	}
	wg.Done()
}

func pinger(start_addrs string, end_addrs string) {
	startip := ip2Long(start_addrs)
	endip := ip2Long(end_addrs)
	addrs := make(chan string, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(addrs); i++ {
		go worker(addrs, &wg)
	}
	for i := startip; i <= endip; i++ {
		wg.Add(1)
		i := int64(i)
		addrs <- back2IP(i)
	}
	wg.Wait()
	close(addrs)
}
