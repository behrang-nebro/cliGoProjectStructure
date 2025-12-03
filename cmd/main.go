package main

import (
	"flag"
	"fmt"

	"github.com/behrang-nebro/cliGoProjectStructure/internals/gomod"
	"github.com/behrang-nebro/cliGoProjectStructure/internals/makedirectory"
)

func main() {
	fmt.Println("hello world")

	projectName := flag.String("projectName", "newProject", "the name of the project to make the structure for")
	githubUsername := flag.String("githubUsername", "notSpecified", "the name of the github username")

	internalsBool := flag.Bool("intenal", true, "should I make a internals directory?")
	pkgBool := flag.Bool("pkg", true, "should I make a pkg directory?")
	apiBool := flag.Bool("api", true, "should I make a api directory?")
	configsBool := flag.Bool("configs", true, "should I make a configs directory?")
	deploymentsBool := flag.Bool("deployments", true, "should I make a deployments directory?")
	migrationsBool := flag.Bool("migrations", true, "should I make a migrations directory?")
	scriptsBool := flag.Bool("scripts", true, "should I make a scripts directory?")
	testBool := flag.Bool("test", true, "should I make a test directory?")
	webBool := flag.Bool("web", true, "should I make a web directory?")

	flag.Parse()

	makedirectory.MkRootDir(*projectName)

	if *internalsBool {
		makedirectory.MkSubDir(*projectName, "internals")
	}

	if *pkgBool {
		makedirectory.MkSubDir(*projectName, "pkg")
	}

	if *apiBool {
		makedirectory.MkSubDir(*projectName, "api")
	}

	if *configsBool {
		makedirectory.MkSubDir(*projectName, "configs")
	}

	if *deploymentsBool {
		makedirectory.MkSubDir(*projectName, "deployments")
	}

	if *migrationsBool {
		makedirectory.MkSubDir(*projectName, "migrations")
	}

	if *scriptsBool {
		makedirectory.MkSubDir(*projectName, "scripts")
	}

	if *testBool {
		makedirectory.MkSubDir(*projectName, "test")
	}

	if *webBool {
		makedirectory.MkSubDir(*projectName, "web")
	}

	gomod.ExecGoMod(*githubUsername, *projectName)
}
