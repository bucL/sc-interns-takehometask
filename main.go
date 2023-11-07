package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {

	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	
	

	// Uncomment this and comment out the below input if you want to print out all of the pagination at once. 
	// Make sure you recomment this out if you want to use the commandLine to enter the next token and check functionality
	// as shown in the readme. Also go to folders_pagination.go and comment out the last if statement in GenerateMap. 
	// It's basically just a starter for the commandline input version. 
	/*
	output:= folders.GenerateMap(res.Folders)
	for k, v:= range output {
		fmt.Println("current key is",k)
		folders.PrettyPrint(v.Data)
		fmt.Println("next key is ", v.NextToken)
	}
	*/
	/* 
	use the following piece of code if you want to use the commandline to enter the next token and check functionality
	comment this out and uncomment the above if you want to print all of them at once
	to check that the last one doesn't have a token search for 'ultimate-pandemic' in the terminal output
	the line that says next key is... shouldn't have anything. */
	
	folders.GenerateMap(res.Folders)
	var input string
	for input != "exit" {
		fmt.Println("Enter a token (to exit simply type 'exit')")
		fmt.Scan(&input)
		if input != "exit" {
			var nextFolderSet = folders.Pagination(input)
			folders.PrettyPrint(nextFolderSet.Data)
			fmt.Println(nextFolderSet.NextToken)
		}
	}
}
