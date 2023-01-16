package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	host := "1.1.1.1"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	fmt.Printf("Finding MTU max:\n")
	fmt.Printf("================\n")
	min := 68
	max := 9000

	for min <= max {
		mtu := (min + max) / 2
		//fmt.Printf("test %d\n", mtu)
		var cmd *exec.Cmd
		if runtime.GOOS != "windows" {
			cmd = exec.Command("ping", "-W", "3", "-c", "1", "-M", "do", "-s", strconv.Itoa(mtu), host)
		} else {
			cmd = exec.Command("ping", "-w", "3", "-f", "-n", "1", "-l", strconv.Itoa(mtu), host)
		}
		err := cmd.Run()
		if err == nil {
			min = mtu + 1
			fmt.Printf("+")
		} else {
			max = mtu - 1
			fmt.Printf("-")
		}
	}
	mtu := max + 28
	fmt.Printf("\nMaximum MTU value for %s: %d\n", host, mtu)
	if runtime.GOOS == "windows" && len(os.Args) == 1 {
		fmt.Println("Appuyez sur Entrée pour continuer...")
		fmt.Scanln()
	}
}
