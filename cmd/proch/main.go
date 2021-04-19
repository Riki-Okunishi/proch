package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"github.com/Riki-Okunishi/proch"
)

func main() {
	// change encoding from Shift-JIS to UTF-8
	chcp := exec.Command("cmd", "/C", "chcp", "65001")
	chcp.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := chcp.Run()
	if err != nil {
		fmt.Printf("Error: Failed to change encoding to UTF-8 by executing 'chcp 65001'\n\t%s\n\n", err)
		os.Exit(1)
	}

	pc := proch.NewProxyChanger()
	pc.Run()

}