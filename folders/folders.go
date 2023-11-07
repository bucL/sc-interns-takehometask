package folders
/*
import (
	"fmt"
	"github.com/gofrs/uuid"

)*/

/* ORIGINAL CODE + Suggested Improvements/Comments + Function:
*/
/*
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	/* 
	f1 and fs are variables that aren't required for the function to work. err here can be used to return an
	error in the output/terminal if one occurs using an if statement at the return. 
	
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	// this statement creates a new 'slice' (aka python list) of the 'Folder' type
	f := []Folder{}

	// Store a map of the output of the FetchAllFoldersByOrgID function in the 'r' variable and ignore 
	// any errors that occur
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	
	// this for loop doesn't work and returns an error because the var 'k' isn't used anywhere
	// in the improved code both this for loop and the one below can be combined into one to reduce time complexity. 
	for k, v := range r {
		f = append(f, *v)
	}
	
	//Creates a new empty 'slice'(aka python list) of pointers to the Folder type stored in 'fp' 
	var fp []*Folder
	
	/* 
	Again this for loop doesn't work in the original code as the variable k1 isn't used anywhere. 
	This can be combined with the previous for loop to achieve the same result which is creating a slice of pointers
	to Folders with the specified orgID.  
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	
	/*
	These lines of code assign the results of the FetchAllFolders to the FetchFolderResponse type and return the result
	however this is done without error handling and will always result in a nil error. 
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}
 

/* 
This function works as intended however it does not have proper error handling as it will always return nil 
even if one does occur. 
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
}*/

// THE FUNCTIONS ARE COMMENTED OUT BECAUSE THEY WERE COPIED AND PLACED IN THE PAGINATION FILE 

// UPDATED CODE with improvements implemented. 
/*
 GetAllFolders retrieves a list of folders based on the provided FetchFolderRequest.
 It calls the FetchAllFoldersByOrgID function to fetch folder data, and then
 constructs a FetchFolderResponse containing the list of folders.

 Parameters:
   req: A pointer to a struct of type FetchFolderRequest containing the organization ID for filtering.

 Returns:
   - A pointer to a struct of type FetchFolderResponse containing the list of folders.
   - An error if any issues occur during the retrieval process including calling the
	   FetchAllFoldersByOrgID function
*/
/*
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
    // Call the FetchAllFoldersByOrgID function to retrieve a map r.
    var err error
	r, err := FetchAllFoldersByOrgID(req.OrgID)

	// if an error occurs return the error from FetchAllFoldersByOrgID
	if err != nil{
		return nil, err
	}
    // Initialize an empty slice of pointers to the Folder type
    var fp []*Folder

    // Loop over the values in the map r and create a slice of pointers
    for _, v := range r {
        fp = append(fp, v)
    }
    // Create a FetchFolderResponse struct with the Folders field set to fp
    ffr := &FetchFolderResponse{Folders: fp}
    // Return the FetchFolderResponse result that is of type FetchFolderResponse. 
	return ffr, err
}
*/

/*
FetchAllFoldersByOrgId retrives a slice of pointers to all Folders that contain the specified orgID by calling the
GetSampleData() function and then looping over it as a range and appending to the resFolder empty slice. 
Parameters: 
- orgID: uuid used to filter folders by specific organisation. 
Returns: 
- resFolder: a slice (aka python list) of pointers to Folder/s that match the orgID. 
- err: if an error occurs return an error statement in the output console. 
*/
/*
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	// Added error handling
	if len(resFolder) == 0 {
		return nil, fmt.Errorf("No folders found for orgID: %s", orgID)
	}
	return resFolder, nil
}
*/