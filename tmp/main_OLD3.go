package main

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
	"log"
)

func main() {

// Declare Variables
	var gheUrl string

// Check OS type
	if runtime.GOOS == "darwin" {
		fmt.Println("You are running MacOS")
	} else {
		fmt.Println("You are running OS other than MacOS")
	}
// Check git exists
	gitPath, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("git is currently not available on this system")
	}
	fmt.Printf("git is available at %s\n", gitPath)

// Below logic is to check number of arguments are correctly passed
	if len(os.Args[1:]) != 4 {
		msg := `This binarys needs 4 arguments - Package Directory, GHE Organisation, Reponame
		And Commit Message
		"Gitcreate <directory> <GHE Org> <Repo> <Commit message>"
		
		`
		fmt.Println(msg)
		os.Exit(30)
	} 

// Grab inputs for processing 
	packageDir	:=	os.Args[1]
	gheOrg		:=	os.Args[2]
	repoName	:=	os.Args[3]
	commitMsg	:=	os.Args[4]
	gheUrl		=	"git@github.com:" + gheOrg + "/" + repoName + ".git" 

	fmt.Println("")
	fmt.Printf("PackageDirectory:: %v \n", packageDir)
	fmt.Printf("GHE Organisation:: %v \n", gheOrg)
	fmt.Printf("Repo Name:: %v \n", repoName)
	fmt.Printf("commitMessage:: %v \n", commitMsg)
	fmt.Printf("GHE URL:: %v \n", gheUrl)

	fmt.Printf("gitclone %s \n", gitClone(gitPath, gheUrl, packageDir))
	fmt.Printf("gitcommit %s \n", gitCommit(gitPath, packageDir, commitMsg))
	//fmt.Printf("gitpush %s \n", gitPush(gitPath))
}

func gitClone(path, gheurl string, directory string) string {
	cmd := exec.Command(path, "clone", gheurl, directory, "--recursive")
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%v \n", string(stdoutStderr))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s \n", err)
	}
	return string(stdoutStderr)
}

func gitCommit(path, directory, msg string) (string) {
	cmd := exec.Command(path, "add --all", directory)
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%v \n", string(stdoutStderr))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s \n", err)
	}
	return string(stdoutStderr)
}

/*
func gitPush(path string) string {
	cmd: exec.Command(path, "push -u origin master")
	stdoutStderr, err := cmd.CombinedOutput()
	
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s \n", err)
	}
	return string(stdoutStderr)
}
*/
