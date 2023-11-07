package folders

import (
	"fmt"
	"github.com/gofrs/uuid"
)
// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
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
/*
FetchAllFoldersByOrgId retrives a slice of pointers to all Folders that contain the specified orgID by calling the
GetSampleData() function and then looping over it as a range and appending to the resFolder empty slice. 
Parameters: 
- orgID: uuid used to filter folders by specific organisation. 
Returns: 
- resFolder: a slice (aka python list) of pointers to Folder/s that match the orgID. 
- err: if an error occurs return an error statement in the output console. 
*/
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


type PaginationOutput struct {
	Data []*Folder
	Token string
}

func Paginate(data []*Folder, token string) PaginationOutput {
	
	return PaginationOutput{}
}