package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/Riki-Okunishi/proch"
)

func main() {
	// change encoding from Shift-JIS to UTF-8
	err := exec.Command("cmd", "/C", "chcp", "65001").Run()
	if err != nil {
		fmt.Printf("Error: Failed to change encoding to UTF-8 by executing 'chcp 65001'\n\t%s\n\n", err)
		os.Exit(1)
	}

	pc := proch.NewProxyChanger()
	pc.Run()

}