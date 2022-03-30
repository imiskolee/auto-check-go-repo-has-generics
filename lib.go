package auto_check_go_repo_has_generics

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
)

var findGoVersionRegex = regexp.MustCompile("go1\\.(\\d+)")

func checkGoVersion() bool {
	buffer := bytes.NewBufferString("")
	cmd := exec.Command("go", "version")
	cmd.Stdout = buffer
	if err := cmd.Run(); err != nil {
		return true
	}
	subs := findGoVersionRegex.FindSubmatch(buffer.Bytes())
	if len(subs) < 2 {
		return true
	}
	minorVersion, _ := strconv.ParseInt(string(subs[1]), 10, 64)
	if minorVersion >= 18 {
		panic("PLEASE DO NOT USING GO 1.18+ NOW")
	}
	return true
}

func init() {
	checkGoVersion()
}
