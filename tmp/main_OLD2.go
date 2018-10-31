package main

import (
	"os"
	"fmt"
	"os/exec"
	"runtime"
	"log"
//	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

func main() {

// Variables Declared
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


// Grab inputs for processing 
	CheckArgs("<packageDir>", "<gheOrg>", "<repoName>", "<commigMsg>")
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


// Gitclone the repository 

	// Clone the given repository to the given packageDir
	Info("git status")

/*
	Info("git clone %s %s --recursive", gheUrl, packageDir)

	r, err := git.PlainClone(packageDir, false, &git.CloneOptions{
		URL:               gheUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
*/


}

