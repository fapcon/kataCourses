package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println(ExecBin("ls", "-la"))
	fmt.Println(ExecBin("nonexistent-binary"))
}

func ExecBin(binPath string, args ...string) string {
	// Ваш код здесь

	cmd := exec.Command(binPath, args...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "Error executing binary:" + err.Error()
	}
	s := string(b)
	return s

}
