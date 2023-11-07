package folders_test

import (
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"

	// Unused import
	//"github.com/stretchr/testify/assert"
	"github.com/gofrs/uuid"
)

/*
Based on the instruction this set of unit tests assume that the FetchAllFoldersByOrgId function is 100% working and therefore
doesn't have a mock FetchAllFoldersByOrgId slice (aka python list) of pointers.
However note that I have updated the FetchAllFoldersByOrgID function to include some sort of error handling
as this wasn't present in the original code and this is used in Test Case 2
*/
func Test_GetAllFolders(t *testing.T) {
	/*
		Test Case 1: Program Succesfully executes when provided with a valid orgID
		To run multiple tests change OrgID to different value that exists in sample.json to confirm it works
	*/
	t.Run("testDefaultSuccess", func(t *testing.T) {
		// Setup for OrgID adapted from main.go
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}
		res, err := folders.GetAllFolders(req)
		if err != nil {
			t.Errorf("Expected no error but instead got: %d", err)
			return
		}
		if res == nil {
			t.Errorf("Expected to have a value for res but instead got nil")
			return
		}
		fmt.Println("Successfully Passed Test 1")
		return
	})

	// Test Case 2: Function returns an error message when provided with an invalid orgID
	t.Run("testErrorMsg", func(t *testing.T) {
		// use a uuid that doesn't exist.
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17b"),
		}
		_, err := folders.GetAllFolders(req)
		if err == nil {
			t.Errorf("Expected an error message but got none")
			return
		}
		if err != nil {
			fmt.Println("Successfully Passed Test 2")
			return
		}

	})

	/*
		Test Case 3: Check if program can successfully handle large amounts of data In this case the default
		orgID has 666 instances in the sample.json and therefore should have a length of 666
	*/
	t.Run("testBulk", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}
		res, err := folders.GetAllFolders(req)
		if err != nil {
			t.Errorf("Expected no error instead encountered: %d", err)
			return
		}
		if len(res.Folders) != 666 {
			t.Errorf("Expected length of output to be 666 instead got length: %d ", len(res.Folders))
			return
		} else {
			fmt.Println("Successfully Passed Test 3")
			return
		}

	})

}
