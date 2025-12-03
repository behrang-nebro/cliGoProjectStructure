package gitrepo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecGit(userName, projectName string) {
	cmdDIr := filepath.Join("./", projectName)

	cmd := exec.Command("git", "init")
	cmd.Dir = cmdDIr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("git init failed: %v", err)
		return
	}

	fmt.Printf("git init was executed succesfuly\n")
	fmt.Printf("################################\n")

	createGitignore(cmdDIr)

}

func createGitignore(projectDir string) error {
	gitignoreContent := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE files
.vscode/
.idea/

# OS generated files
.DS_Store
Thumbs.db

# Environment variables
.env
.env.local

# Logs
*.log`

	// Write to file
	gitignorePath := filepath.Join(projectDir, ".gitignore")
	return os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644)
}
