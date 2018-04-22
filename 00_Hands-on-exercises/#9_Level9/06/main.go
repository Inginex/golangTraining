package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	gover := runtime.Version()
	osver := runtime.GOOS
	osarch := runtime.GOARCH
	gid := os.Getpid()
	gwd, _ := os.Getwd()
	hstn, _ := os.Hostname()
	env := os.Environ()

	fmt.Println("NUM CPU: ", cpu)
	fmt.Println("GO Version: ", gover)
	fmt.Println("OS SYSTEM: ", osver)
	fmt.Println("OS ARCH: ", osarch)
	fmt.Println("GID: ", gid)
	fmt.Println("Getwd: ", gwd)
	fmt.Println("Hostname: ", hstn)
	fmt.Println("Environment: ", env)
}
