package user

import (
	"testing"
	"time"
)

func TestNew_Validate(t *testing.T) {
	nu := New{
		Email:     "gayshit@gmail.com",
		Username:  "gayshit",
		Firstname: "Gay",
		Lastname:  "Shit",
		BirthDate: time.Now(),
		Password:  "23490239ashljkfdhalv",
	}

	res, _ := nu.Validate()

	if res == nil {
		t.Log("MAP IS NILLL")
	}
	t.Logf("res : %v", res)

}
