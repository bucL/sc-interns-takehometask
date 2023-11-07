package folders

import (
	"fmt"
	"github.com/gofrs/uuid"
	"math/rand"
)

// Used to generate tokens. 
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var output map[string]PaginationOutput


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
    // Call the FetchAllFoldersByOrgID function to retrieve a map r that contains all folders with the req OrgID
    var err error
	r, err := FetchAllFoldersByOrgID(req.OrgID)

	// if an error occurs return the error from FetchAllFoldersByOrgID
	if err != nil{
		return nil, err
	}
    // Initialize an empty slice of pointers to the Folder type
    var fp []*Folder

    //Append the contents of the map r to fp. 
	fp = append(fp, r...)

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
		return nil, fmt.Errorf("no folders found for orgID: %s", orgID)
	}
	return resFolder, nil
}




/*
METHOD USED TO ACHIEVE PAGINATION:

I honestly didn't know about pagination before this but here is my attempt at implementating it into this program. 
The functionality of this is split into 4 different functions. This implementation uses userInput into the commandLine 
to simulate a user clicking a button on a webpage or app for example. 

The main function responsible for all the heavylifting is the GenerateMap function which takes in the output from the GetAllFolders
function and returns a map with unique strings, aka tokens, as keys and a struct of type PaginationOutput which contains the data,
aka pointers to Folders and the nextToken as a string. 

First the GenerateMap function will call the GenerateTokens function which uses the output list of all folders that have the 
required OrgID and generates a number of unique tokens using a randomstring generator that is located in RandStringBytes()
Then it will initatise a map object stored in 'mapped' before running a loop to assign each token to a PaginationOutput struct
which contains a pair of Folders and the nextToken. 'mapped' is then assigned to the global variable 'output' which 
is returned by the function. Then if you are using the commandLine to enter each token 1 by 1 the Pagination function is called
which simulates user interaction and loads the next set of data with the code for prettyprint in main.go 


*/

type PaginationOutput struct {
	Data []*Folder
	NextToken string
}

// Random String Generation Function Taken from https://stackoverflow.com/a/31832326
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}
/* 
GenerateTokens generates a slice of unique tokens based on the provided list of folders.
Each token is created using the RandStringBytes function and represents a pagination marker.
Parameters:
 Data: A slice of pointers to Folder objects from which tokens will be generated.
Returns:
 A slice of unique tokens that can be used for pagination.
The function calculates the number of tokens to generate, which is half the number of folders
in the provided list. It uses the RandStringBytes function to create each token. The generated
tokens can be used to navigate and paginate through the list of folders. */
func GenerateTokens(Data []*Folder) []string {
	var output []string
	numberTokens := len(Data) / 2 - 1
	for x:= 0; x<= numberTokens; x++ {
		output = append(output, (RandStringBytes(5)))
	}
	fmt.Println("total number of tokens: ", len(output))
	return output
}

// GenerateMap generates a map of pagination data based on a list of folders and associated tokens.
//
// Parameters:
//   listOfFolders: A slice of pointers to Folder objects that will be paginated.
//
// Returns:
//   A map where each key is a pagination token and each value is a PaginationOutput struct.
//
// The function first generates pagination tokens using the GenerateTokens function. It then
// constructs a map where each token serves as a key, and the corresponding PaginationOutput struct
// is the associated value. Each PaginationOutput struct contains a Data field with a slice of Folder
// objects for the current page, and a NextToken field representing the token for the next page.
//
// Note: The PrettyPrint function is called for the first PaginationOutput struct in the map if paginating through. 
func GenerateMap(listOfFolders []*Folder) map[string]PaginationOutput {
	tokens:= GenerateTokens(listOfFolders)
	mapped := make(map[string]PaginationOutput)
	for n := 0; n< len(tokens); n++ {
		var x PaginationOutput
		if n+1== len(tokens) {
			x.NextToken = ""
		} else{
			x.NextToken = tokens[n+1]	
		}
		startIndex := n * 2
        endIndex := (n+1)*2
        if endIndex > len(listOfFolders) {
            endIndex = len(listOfFolders)
        }
        x.Data = append(x.Data, listOfFolders[startIndex:endIndex]...)

		mapped[tokens[n]] = x
		// comment this 'if' out if you want to see the entire map at once. 
		if n == 0 {
			PrettyPrint(x.Data)
			fmt.Println("Next Token is", x.NextToken)
		}
	}
	output = mapped
	return output
}

func Pagination(request string) PaginationOutput {
	return output[request]
}