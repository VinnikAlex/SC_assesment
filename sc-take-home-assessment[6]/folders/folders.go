package folders

import (
	"github.com/gofrs/uuid"
)

// Improvements for GetAllFolders() method:
// 1. Variable names are hard to understand regarding the type of data they represent. Clearer and more descriptive naming conventions would be beneficial.
// 2. Variables err,f1,fs are declared but not used? Instead of creating extra variables such as *fp*, existing variable *fs* should be used.
//	  Other variables should not exist if not used.
// 3. Comments throughout the code would further make it more readable - e.g. describing what a method does and what it returns.

// A method that returns all folders by OrgID
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// purely declares a set of variables - *err*, *f1*, *fs*
	var (
		err error
		f1  Folder
		fs  []*Folder
	)

	// declares variables *f* and *r*. *f* is assigned the Folder constructor with an ID, Name, OrgID and Deleted. *r* is assigned the result
	// of the method FetchAllFoldersByOrgID(req.OrgID)
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	// loops over all folders returned by FetchAllFoldersByOrgID(req.OrgID) and appends each folder to *f*. Allocates unique ID, name, org_id, and deleted status to each folder.
	for k, v := range r {
		f = append(f, *v)
	}

	// loops over all folders in *f* and appends to new variable *fp*
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	
	// set variable *ffr* that points FetchFolderResponse
	var ffr *FetchFolderResponse
	// variable *ffr* address - points to FetchFolderResponse *fp*
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// A method that retreives all folders by org_id in sample.json
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// declaring folders and assigning data from sample.json
	folders := GetSampleData()

	resFolder := []*Folder{}
	// loops through each folder and check if OrgId matches the OrdId passed into this method (FetchAllFoldersByOrgID)
	for _, folder := range folders {
		if folder.OrgId == orgID {
			// if OrgId matches, append the folder to new array resFolder
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
