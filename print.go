package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// PrintHyperLink echo link on console.
//
// `echo -e '\e]8;;http://example.com\aThis is a link\e]8;;\a'`
func PrintHyperLink(link, text string) {
	linkMsg := fmt.Sprintf(`\e]8;;%s\a%s\e]8;;\a`, link, text)
	cmd := exec.Command("echo", "-e", linkMsg)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DateToString(t time.Time) string {
	return t.Format("2006-01-02")
}

func Err(msg string, args ...interface{}) {
	errMsg := fmt.Sprintf("ERROR: %s\n", msg)
	fmt.Printf(errMsg, args...)
}

func Info(msg string, args ...interface{}) {
	infoMsg := fmt.Sprintf("INFO: %s\n", msg)
	fmt.Printf(infoMsg, args...)
}

func Print(msg string, args ...interface{}) {
	msg = fmt.Sprintf("%s\n", msg)
	fmt.Printf(msg, args...)
}
