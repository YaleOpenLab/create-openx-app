package core

import (
	"github.com/pkg/errors"

	utils "github.com/Varunram/essentials/utils"
	openx "github.com/YaleOpenLab/openx/database"
)

// Recipient defines the recipient structure
type Recipient struct {
	U *openx.User
	// user related functions are called as an instance directly
	ReceivedSolarProjects       []string
	ReceivedSolarProjectIndices []int
}

// NewRecipient creates and returns a new recipient
func NewRecipient(uname string, pwd string, seedpwd string, Name string) (Recipient, error) {
	var a Recipient
	var err error
	user, err := NewUser(uname, utils.SHA3hash(pwd), seedpwd, Name)
	if err != nil {
		return a, errors.Wrap(err, "failed to retrieve new user")
	}
	a.U = &user
	err = a.Save()
	return a, err
}
