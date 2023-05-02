package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	title := flag.String("title", "Notification", "Notification Title")
	message := flag.String("message", "", "Notification Message")
	sound := flag.String("sound", "default", "Notification Sound")
	flag.Parse()

	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s" sound name "%s"`, *message, *title, *sound))

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
