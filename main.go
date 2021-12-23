package main

import (
	"fmt"
	"time"
	"syscall"
)

var (
	k32 = syscall.NewLazyDLL("kernel32.dll")
	u32 = syscall.NewLazyDLL("user32.dll")

	// Input
	pGetAsyncKeyState = u32.NewProc("GetAsyncKeyState")
)


func IsPressed(key uint32) bool {
	ret, _, _ := pGetAsyncKeyState.Call(uintptr(key))
	return ret != 0
}

func main(){
	fmt.Println("Hold and release CAPS LOCK to measure time")
	wasPressed := false
	startTime := time.Now()

	for {
		if IsPressed(0x14){
			if !wasPressed {
				startTime = time.Now()
				wasPressed = true
			}
		} else {
			if wasPressed {
				endTime := time.Since(startTime).Seconds()
				fmt.Println("Time pressed:", endTime)
				wasPressed = false
			}
		}
	}
}