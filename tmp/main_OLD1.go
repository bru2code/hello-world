package main

import (
	"os"
	"fmt"
//	"runtime"
	"os/exec"
	"log"
)

func main() {
/*
// Debug Statements Block Begin-	
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	args	:= len(os.Args[1:])
	fmt.Printf("Number of Arguments:: %v \n", args)
	fmt.Printf("OS :: %v \n", runtime.GOOS)

	// Prints the arguments as is
	fmt.Println("")
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	// Check OS type
	if runtime.GOOS == "darwin" {
		fmt.Println("You are running MacOS")
	} else {
		fmt.Println("You are running OS other than MacOS")
	}
*/
	// Check git exists
	gitPath, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("git is currently not available on this system")
	}
	fmt.Printf("git is available at %s\n", gitPath)

// Debug Statement Block End	

	var gheUrl string

// Below logic is to check number of arguments are correctly passed
	if len(os.Args[1:]) != 4 {
		fmt.Println("")
		fmt.Println("This binary needs 4 arguments - Package Directory, GHE Organisation, Reponame")
		fmt.Println("And Commit Message")
		fmt.Println("Gitcreate <directory> <GHE Org> <Repo> <Commit message>")
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

	gitStatus := gitClone(gitPath, gheUrl)
	fmt.Printf("gitstatus is %s \n", gitStatus)
        gitCommit := gitCommit	
	copy.copy(packageDir, repoName)
	
/*
	gitClone(gitPath, gheUrl)
	gitAdd()
	gitCommit()
	gitPush()
	
*/

}

func gitClone(path string, gheurl string) string {
	cmd := exec.Command(path, "clone", gheurl)
	stdoutStderr, err := cmd.CombinedOutput()
//	fmt.Printf("CombinedOutput %s\n", stdoutStderr)
	
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s \n", err)
	}
//	fmt.Printf("Combined output : \n %s\n", string(stdoutStderr))
	return string(stdoutStderr)
	//return "This is a test message"
}
