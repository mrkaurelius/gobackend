package database_test

import (
	"testing"

	"github.com/mrkaurelius/gobackend/database"
)

// trivial test
func TestAllUserPostsJSON(t *testing.T) {
	got := database.AllUserPostsJSON()
	if len(got) < 32 {
		t.Error("checksum size ")
	}
	t.Log(string(got))
}

func TestValidateUser(t *testing.T) {
	got := database.ValidateUser("mrk", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4")
	if got != true {
		t.Errorf("error")
	}
}
