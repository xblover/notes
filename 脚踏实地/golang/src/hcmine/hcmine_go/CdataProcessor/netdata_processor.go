package CdataProcessor

/*
//#cgo LDFLAGS: -L ./ -lhcmine -lstdc++ -ldl
//#cgo CFLAGS: -I .

#include <stdlib.h>
#include <stdio.h>
#include "../event_for_go.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

var netQueue = make(chan *C.char, 10000)

/*
get tcp/udp data from libhcmine.so,
then put it to netQueue,
*/
func getNetDataFromC() {
	count := 0
	period := time.Duration(3000 * time.Microsecond)
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		fmt.Print("http count:", count)
		select {
		case <-ticker.C:
		}
		re := C.getNetData()
		if re == nil {
			time.Sleep(100 * time.Microsecond)
			continue
		}
		count++
		netQueue <- re
	}
}

/*
Consumption (消费) queue data,
you cloud write some business code in this func
*/
func analysisNetData() {
	for {
		re := <-netQueue
		res := C.GoString(re) // GoString ???
		logp.Debug("main", "resData: %s", res)
		if re != nil {
			C.free(unsafe.Pointer(re))
			re = nil
		}
	}
}

//
func regularBus() {
	count := 0
	period := 3 * time.Second
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		}
		count++

		c.getTimeOutNetData(2)
		if count == 20 {
			C.getTimeOutNetData(1)
			count = 0
		}
	}
}

func NetDataProcessorStart(conntrackOn int) {
	go regularBus()
	go getNetDataFromC()
	go analysisNetData()
	go C.startConntrack(C.int(conntrackOn)) //C.startConntrack ???
}
