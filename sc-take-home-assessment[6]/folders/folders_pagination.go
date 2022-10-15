package folders

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started

import (
	"github.com/gofrs/uuid"
)


func PagGetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	var (
		err error
		f1  Folder
		fs  []*Folder
	)


	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	
	for k, v := range r {
		f = append(f, *v)
	}


	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	

	var ffr *FetchFolderResponse

	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

func PagFetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {

	folders := GetSampleData()

	resFolder := []*Folder{}

	for _, folder := range folders {
		if folder.OrgId == orgID {
		
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
