package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"runtime"

	"github.com/Riki-Okunishi/proch"
)

func main() {
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("running on Windows. continue executing proch.\n")
	default:
		fmt.Printf("running on OS not supported. exit proch.")
		os.Exit(1)
	}

	// change encoding from Shift-JIS to UTF-8
	chcp := exec.Command("cmd", "/C", "chcp", "65001")
	chcp.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := chcp.Run()
	if err != nil {
		fmt.Printf("Error: Failed to change encoding to UTF-8 by executing 'chcp 65001'\n\t%s\n\n", err)
		os.Exit(1)
	}

	pc := proch.New(proch.NewNetshRunner(), proch.NewRegistryEditor(), proch.NewJsonLoader())
	pc.Run()
}