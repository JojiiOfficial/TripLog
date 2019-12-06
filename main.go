package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	t, err := tail.TailFile("/var/log/fail2ban.log", tail.Config{Follow: true, ReOpen: true, Poll: true})
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
