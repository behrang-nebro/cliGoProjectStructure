package makedirectory

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func MkRootDir(projectName string) {
	err := os.MkdirAll(projectName, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the root directory --> %12s <-- was created\n", projectName)
}

func MkSubDir(projectName, subDir string) {
	path := filepath.Join(projectName, subDir)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" the sub directory --> %12s <-- was created under --> %12s <--\n", subDir, projectName)
}
