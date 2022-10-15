package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/stretchr/testify/assert"
)

	func Test_GetAllFolders(t *testing.T) {
		allFolders:= folders.GetSampleData()

		got:= GetAllFolders(allFolders)

		want:= allFolders

		if(got != want) {
			t.Error("got %q, wanted %q", got, want)
		}
	}