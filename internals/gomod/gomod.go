package gomod

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecGoMod(userName, projectName string) {
	goModPath := filepath.Join("github.com", userName)
	cmdDIr := filepath.Join("./", projectName)

	cmd := exec.Command("go", "mod", "init", goModPath)
	cmd.Dir = cmdDIr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("go mod init failed: %v", err)
		return
	}

	fmt.Printf("go mod init was executed succesfuly\n")

}
